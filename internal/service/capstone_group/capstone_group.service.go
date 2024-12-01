package capstone_group_service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/capstone_group_dto"
	"github.com/api/internal/dto/user_dto"
	"github.com/api/internal/queue"
	context_util "github.com/api/pkg/utils/context"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/thoas/go-funk"
)

type ICapstoneGroupService interface {
	CreateCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.CreateCapstoneGroupInput) (*capstone_group_dto.CapstoneGroupOutput, error)
	UpdateCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.UpdateCapstoneGroupInput) error
	InviteMentorToCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.InviteMentorToCapstoneGroupInput) error
	ResponseInviteMentorToCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.ResponseInviteMentorToCapstoneGroupInput) error
	GetCapstoneGroup(ctx *gin.Context, id int) (*capstone_group_dto.CapstoneGroupWithTotalMemberOutput, error)
	GetListCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.GetListCapstoneGroupInput) (*capstone_group_dto.ListCapstoneGroupWithTotalMemberOutput, error)
	GetCurrentListCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.GetCurrentListCapstoneGroupInput) (*[]capstone_group_dto.CapstoneGroupWithTotalMemberOutput, error)
	GetMentorAndListMemberCapstoneGroup(ctx *gin.Context, id int64) (*capstone_group_dto.MentorAndListMemberCapstoneGroupOutput, error)
	GetListInvitationMentorCapstoneGroups(ctx *gin.Context, input *capstone_group_dto.GetListInviteMentorToCapstoneGroupInput) (*capstone_group_dto.ListInvitationMentorCapstoneGroupOutput, error)
	GetListStudentHaveCapstoneGroup(ctx *gin.Context, semesterID int64) (*[]*user_dto.StudentOutput, error)
	UpdateCapstoneGroupStudent(ctx *gin.Context, input *capstone_group_dto.UpdateCapstoneGroupStudentInput) error
	UpdateCapstoneGroupReportDocument(ctx *gin.Context, input *capstone_group_dto.UpdateCapstoneGroupReportDocumentInput) error
	DeleteCapstoneGroupReportDocument(ctx *gin.Context, reportID int64, capstoneGroupID int64) error
	GetCapstoneGroupReportDocument(ctx *gin.Context, id int64) (*capstone_group_dto.ReportDocumentOutput, error)
	GetCapstoneGroupReportDocuments(ctx *gin.Context, capstoneGroupID int64) ([]capstone_group_dto.ReportDocumentOutput, error)
	MentorUpdateStudentScoreForReportDocument(ctx *gin.Context, input *capstone_group_dto.MentorUpdateStudentScoreForReportDocument) error
	GetReportDocumentStudentsScore(ctx *gin.Context, reportDocumentID int64) ([]capstone_group_dto.ReportDocumentStudentScoreOutput, error)
	AdminUpdateStudentScoreForReportDocument(ctx *gin.Context, input *capstone_group_dto.AdminUpdateStudentScoreForReportDocument) error
	CommentReport(ctx *gin.Context, input *capstone_group_dto.CommentReportInput) error
	UpdateCommentReport(ctx *gin.Context, input *capstone_group_dto.UpdateCommentReportInput) error
	DeleteCommentReport(ctx *gin.Context, input *capstone_group_dto.DeleteCommentReportInput) error
	GetListReportComments(ctx *gin.Context, input *capstone_group_dto.GetListCommentReportInput) (*[]*capstone_group_dto.ReportCommentWithUserInfoOutput, error)
}

type capstoneGroupService struct {
	emailInviteMentorToCapstoneGroupPublisher queue.IBasePublisher[queue.InviteMentorToCapstoneGroupMessage]
}

func NewCapstoneGroupService(emailInviteMentorToCapstoneGroupPublisher queue.IBasePublisher[queue.InviteMentorToCapstoneGroupMessage]) ICapstoneGroupService {
	return &capstoneGroupService{
		emailInviteMentorToCapstoneGroupPublisher: emailInviteMentorToCapstoneGroupPublisher,
	}
}

func (cgs *capstoneGroupService) CreateCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.CreateCapstoneGroupInput) (*capstone_group_dto.CapstoneGroupOutput, error) {
	if !funk.Contains(input.StudentIDs, input.LeaderID) {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidLeader,
		}))
	}

	if len(input.StudentIDs) > constant.MaxTotalMemberInGroup || len(input.StudentIDs) < constant.MinTotalMemberInGroup {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidTotalMemberInGroup,
		}))
	}

	var memberGroups []model.Student
	if err := global.Db.Model(model.Student{}).Joins("User").Where("students.id IN ?", input.StudentIDs).Find(&memberGroups).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	if len(memberGroups) != len(input.StudentIDs) {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	var memberExistInAnotherGroup []model.StudentCapstoneGroup
	if err := global.Db.Model(model.StudentCapstoneGroup{}).Where("student_id IN ? AND semester_id = ?", input.StudentIDs, input.SemesterID).Find(&memberExistInAnotherGroup).Error; err == nil {
		if len(memberExistInAnotherGroup) > 0 {
			memberIDExistInAnotherGroup := funk.Map(memberExistInAnotherGroup, func(member model.StudentCapstoneGroup) int64 {
				return member.StudentID
			})

			studentExist := funk.Filter(memberGroups, func(student model.Student) bool {
				return funk.ContainsInt64((memberIDExistInAnotherGroup).([]int64), student.ID)
			})

			memberNameArr := funk.Map(studentExist, func(student model.Student) string {
				return student.User.Name
			}).([]string)

			memberNames := strings.Join(memberNameArr, ", ")

			return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.MemberExistInAnotherGroup,
				TemplateData: map[string]interface{}{
					"MemberNames": memberNames,
				},
			}))
		}
	}

	var major model.Major
	if err := global.Db.Model(model.Major{}).Where("id = ?", input.MajorID).First(&major).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.MajorNotFound,
		}))
	}

	var semester model.Semester
	if err := global.Db.Model(model.Semester{}).Where("id = ?", input.SemesterID).First(&semester).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.SemesterNotFound,
		}))
	}

	data := model.CapstoneGroup{
		NameGroup:  input.NameGroup,
		MajorID:    input.MajorID,
		SemesterID: input.SemesterID,
		LeaderID:   input.LeaderID,
		Status:     constant.CapstoneGroupStatus.ReviewingTopic,
	}

	tx := global.Db.Begin()
	if tx.Error != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		fmt.Println("Failed to begin a transaction, Error: ", tx.Error)
		return nil, errors.New(message)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Model(model.CapstoneGroup{}).Create(&data).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	studentCapstoneGroups := make([]model.StudentCapstoneGroup, 0, len(input.StudentIDs))
	for _, id := range input.StudentIDs {
		studentCapstoneGroups = append(studentCapstoneGroups, model.StudentCapstoneGroup{
			StudentID:       id,
			SemesterID:      input.SemesterID,
			CapstoneGroupID: data.ID,
		})
	}
	if err := tx.Model(model.StudentCapstoneGroup{}).Save(&studentCapstoneGroups).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for index, reportDocumentTypeMapping := range constant.ReportDocumentTypeMappings {
		reportDocument := model.ReportDocument{
			Name:               constant.NameReportDocumentMappings[index],
			FileIDs:            []string{},
			CapstoneGroupID:    data.ID,
			TypeReport:         reportDocumentTypeMapping,
			MentorReviewStatus: constant.MentorReviewStatusReport.Reviewing,
		}

		if err := tx.Model(model.ReportDocument{}).Create(&reportDocument).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		var studentScores []model.ReportDocumentStudentScore
		for _, studentID := range input.StudentIDs {
			studentScores = append(studentScores, model.ReportDocumentStudentScore{
				StudentID:        studentID,
				ReportDocumentID: reportDocument.ID,
				Score:            nil,
			})
		}

		if err := tx.Model(model.ReportDocumentStudentScore{}).Save(&studentScores).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		return nil, errors.New(message)
	}

	var capstoneGroup model.CapstoneGroup
	if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", data.ID).First(&capstoneGroup).Error; err != nil {
		return nil, err
	}

	capstoneGroupOutput := capstone_group_dto.ToCapstoneGroupOutput(&capstoneGroup)

	return &capstoneGroupOutput, nil
}

func (cgs *capstoneGroupService) UpdateCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.UpdateCapstoneGroupInput) error {
	currentUser := context_util.GetUserContext(ctx)
	if currentUser == nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.PermissionDenied,
		}))
	}

	var currentStudent model.Student
	if err := global.Db.Model(model.Student{}).Where("user_id = ?", currentUser.ID).First(&currentStudent).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	var capstoneGroup model.CapstoneGroup
	if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", input.ID).First(&capstoneGroup).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}

	if capstoneGroup.LeaderID != currentStudent.ID {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.PermissionDenied,
		}))
	}

	if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", input.ID).Updates(&model.CapstoneGroup{
		NameGroup: input.NameGroup,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (cgs *capstoneGroupService) GetCapstoneGroup(ctx *gin.Context, id int) (*capstone_group_dto.CapstoneGroupWithTotalMemberOutput, error) {
	var capstoneGroupWithTotalMember model.CapstoneGroupWithTotalMember
	if err := global.Db.Model(model.CapstoneGroup{}).
		Select("capstone_groups.*, COUNT(sg.student_id) AS total_members").
		Joins("LEFT JOIN student_capstone_groups AS sg ON sg.capstone_group_id = capstone_groups.id").
		Where("capstone_groups.id = ?", id).
		Group("capstone_groups.id").
		First(&capstoneGroupWithTotalMember).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}

	capstoneGroupOutput := capstone_group_dto.ToCapstoneGroupWithTotalMemberOutput(&capstoneGroupWithTotalMember)
	return &capstoneGroupOutput, nil
}

func (cgs *capstoneGroupService) GetListCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.GetListCapstoneGroupInput) (*capstone_group_dto.ListCapstoneGroupWithTotalMemberOutput, error) {
	var total int64
	var items []model.CapstoneGroupWithTotalMember

	query := global.Db.Model(model.CapstoneGroup{}).
		Select("capstone_groups.*, COUNT(sg.student_id) AS total_members").
		Joins("LEFT JOIN student_capstone_groups AS sg ON sg.capstone_group_id = capstone_groups.id").
		Group("capstone_groups.id")

	queryTotal := global.Db.Model(model.CapstoneGroup{})

	if input.SemesterID != 0 {
		query = query.Where("capstone_groups.semester_id = ?", input.SemesterID)
		queryTotal = query.Where("capstone_groups.semester_id = ?", input.SemesterID)
	}

	if input.Status != nil {
		query = query.Where("capstone_groups.status = ?", *input.Status)
		queryTotal = query.Where("capstone_groups.status = ?", *input.Status)
	}

	if err := queryTotal.Count(&total).Error; err != nil {
		return nil, err
	}

	if err := query.
		Limit(int(input.Limit)).
		Offset(int(input.Offset)).
		Find(&items).Error; err != nil {
		return nil, err
	}

	itemsCapstoneGroupOutput := make([]capstone_group_dto.CapstoneGroupWithTotalMemberOutput, len(items))
	for i, item := range items {
		itemsCapstoneGroupOutput[i] = capstone_group_dto.ToCapstoneGroupWithTotalMemberOutput(&item)
	}

	return &capstone_group_dto.ListCapstoneGroupWithTotalMemberOutput{
		Meta: dto.MetaPagination{
			CurrentPage: int(input.Page),
			Total:       int(total),
		},
		Items: itemsCapstoneGroupOutput,
	}, nil
}

func (cgs *capstoneGroupService) GetMentorAndListMemberCapstoneGroup(ctx *gin.Context, id int64) (*capstone_group_dto.MentorAndListMemberCapstoneGroupOutput, error) {
	var capstoneGroup model.CapstoneGroup
	if err := global.Db.Model(model.CapstoneGroup{}).Joins("Mentor.User").Where("capstone_groups.id = ?", id).First(&capstoneGroup).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}
	mentor := capstoneGroup.Mentor
	var mentorOutput *user_dto.TeacherOutput
	mentorOutput = nil
	if mentor != nil {
		mentorOutput = user_dto.ToTeacherOutput(mentor)
	}

	var membersCapstoneGroup []model.StudentCapstoneGroup
	if err := global.Db.Model(model.StudentCapstoneGroup{}).Joins("Student.User").
		Where("student_capstone_groups.capstone_group_id = ?", id).
		Find(&membersCapstoneGroup).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}

	var output capstone_group_dto.MentorAndListMemberCapstoneGroupOutput
	output.LeaderID = capstoneGroup.LeaderID
	output.Mentor = mentorOutput
	output.Members = make([]*user_dto.StudentOutput, len(membersCapstoneGroup))

	for index, memberCapstoneGroup := range membersCapstoneGroup {
		member := memberCapstoneGroup.Student
		memberOutput := user_dto.ToStudentOutput(&member)
		output.Members[index] = memberOutput
	}

	return &output, nil
}

func (cgs *capstoneGroupService) GetListStudentHaveCapstoneGroup(ctx *gin.Context, semesterID int64) (*[]*user_dto.StudentOutput, error) {
	var studentCapstoneGroups []model.StudentCapstoneGroup
	query := global.Db.Model(model.StudentCapstoneGroup{}).
		Joins("Student.User").
		Where("student_capstone_groups.semester_id = ?", semesterID)

	if err := query.Find(&studentCapstoneGroups).Error; err != nil {
		return nil, err
	}

	items := make([]*user_dto.StudentOutput, len(studentCapstoneGroups))
	for i, studentCapstoneGroup := range studentCapstoneGroups {
		items[i] = user_dto.ToStudentOutput(&studentCapstoneGroup.Student)
	}

	return &items, nil
}

func (cgs *capstoneGroupService) UpdateCapstoneGroupStudent(ctx *gin.Context, input *capstone_group_dto.UpdateCapstoneGroupStudentInput) error {
	var capstoneGroup model.CapstoneGroup
	if err := global.Db.Model(model.CapstoneGroup{}).Where("id = ?", input.ID).First(&capstoneGroup).Error; err != nil {
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.CapstoneGroupNotFound,
		}))
	}

	var capstoneGroupReportDocuments []model.ReportDocument
	if err := global.Db.Model(model.ReportDocument{}).
		Where("capstone_group_id = ?", input.ID).
		Find(&capstoneGroupReportDocuments).Error; err != nil {
		global.Db.Rollback()
		return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.FailedToRemoveStudent,
		}))
	}

	reportDocumentIDs := funk.Map(capstoneGroupReportDocuments, func(capstoneGroupReportDocument model.ReportDocument) int64 {
		return capstoneGroupReportDocument.ID
	}).([]int64)

	tx := global.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, studentID := range input.StudentIDs {
		var existingStudentGroup model.StudentCapstoneGroup
		err := tx.Model(model.StudentCapstoneGroup{}).
			Where("student_id = ? AND semester_id = ? AND capstone_group_id != ?", studentID, capstoneGroup.SemesterID, input.ID).
			First(&existingStudentGroup).Error
		if err == nil {
			tx.Rollback()
			return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.StudentAlreadyInAnotherGroup,
			}))
		}

		var studentCapstoneGroup model.StudentCapstoneGroup
		err = tx.Model(model.StudentCapstoneGroup{}).
			Where("capstone_group_id = ? AND student_id = ?", input.ID, studentID).
			First(&studentCapstoneGroup).Error
		if err == nil {
			if studentID == capstoneGroup.LeaderID {
				tx.Rollback()
				return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: constant.MessageI18nId.CannotRemoveLeader,
				}))
			}

			if err := tx.Delete(&studentCapstoneGroup).Error; err != nil {
				tx.Rollback()
				return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: constant.MessageI18nId.FailedToRemoveStudent,
				}))
			}

			if err := tx.Exec(`
				DELETE FROM "report_document_student_scores"
				WHERE "student_id" = ? AND "report_document_id" IN ?`,
				studentID,
				reportDocumentIDs,
			).Error; err != nil {
				tx.Rollback()
				return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: constant.MessageI18nId.FailedToRemoveStudent,
				}))
			}

			continue
		}

		newStudentCapstoneGroup := model.StudentCapstoneGroup{
			CapstoneGroupID: input.ID,
			StudentID:       studentID,
			SemesterID:      capstoneGroup.SemesterID,
		}

		if err := tx.Create(&newStudentCapstoneGroup).Error; err != nil {
			tx.Rollback()
			return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.FailedToAddStudent,
			}))
		}

		var studentScores []model.ReportDocumentStudentScore
		for _, reportDocumentID := range reportDocumentIDs {
			studentScores = append(studentScores, model.ReportDocumentStudentScore{
				ReportDocumentID: reportDocumentID,
				StudentID:        studentID,
				Score:            nil,
			})
		}

		if err := tx.Model(model.ReportDocumentStudentScore{}).Save(&studentScores).Error; err != nil {
			tx.Rollback()
			return errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.FailedToRemoveStudent,
			}))
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (cgs *capstoneGroupService) getCurrentStudent(ctx *gin.Context) (*model.Student, error) {
	currentUser := context_util.GetUserContext(ctx)
	if currentUser == nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	var currentStudent model.Student
	if err := global.Db.Model(model.Student{}).Where("user_id = ?", currentUser.ID).First(&currentStudent).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	currentStudent.User = model.User{
		ID:       currentUser.ID,
		Email:    currentUser.Email,
		Name:     currentUser.Name,
		UserType: currentUser.UserType,
	}

	return &currentStudent, nil
}

func (cgs *capstoneGroupService) GetCurrentListCapstoneGroup(ctx *gin.Context, input *capstone_group_dto.GetCurrentListCapstoneGroupInput) (*[]capstone_group_dto.CapstoneGroupWithTotalMemberOutput, error) {
	currentStudent, currentTeacher, err := cgs.getCurrentStudentOrTeacher(ctx)

	if err != nil {
		return nil, err
	}
	var items []model.CapstoneGroupWithTotalMember

	query := global.Db.Model(model.CapstoneGroup{}).
		Select("capstone_groups.*, COUNT(sg.student_id) AS total_members").
		Joins("LEFT JOIN student_capstone_groups AS sg ON sg.capstone_group_id = capstone_groups.id").
		Group("capstone_groups.id").
		Where("capstone_groups.semester_id = ?", input.SemesterID)

	if currentTeacher != nil {
		query.Where("capstone_groups.mentor_id = ?", currentTeacher.ID)
	}

	if currentStudent != nil {
		query.Where("sg.student_id = ?", currentStudent.ID)
	}

	if err := query.Find(&items).Error; err != nil {
		return nil, err
	}

	output := make([]capstone_group_dto.CapstoneGroupWithTotalMemberOutput, len(items))
	for i, item := range items {
		output[i] = capstone_group_dto.ToCapstoneGroupWithTotalMemberOutput(&item)
	}

	return &output, nil
}

func (cgs *capstoneGroupService) getCurrentTeacher(ctx *gin.Context) (*model.Teacher, error) {
	currentUser := context_util.GetUserContext(ctx)
	if currentUser == nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	var currentTeacher model.Teacher
	if err := global.Db.Model(model.Teacher{}).Where("user_id = ?", currentUser.ID).First(&currentTeacher).Error; err != nil {
		return nil, errors.New(global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.UserNotFound,
		}))
	}

	currentTeacher.User = model.User{
		ID:       currentUser.ID,
		Email:    currentUser.Email,
		Name:     currentUser.Name,
		UserType: currentUser.UserType,
	}

	return &currentTeacher, nil
}
