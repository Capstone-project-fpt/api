package capstone_group_dto

import (
	"time"

	"github.com/api/database/model"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/user_dto"
)

type ReportDocumentStudentScoreOutput struct {
	ID               int64                   `json:"id"`
	StudentID        int64                   `json:"student_id"`
	Student          *user_dto.StudentOutput `json:"student"`
	Score            *float64                `json:"score"`
	ReportDocumentID int64                   `json:"report_document_id"`
	CreatedAt        time.Time               `json:"created_at"`
	UpdatedAt        time.Time               `json:"updated_at"`
}

func ToReportDocumentStudentScoreOutput(studentScore *model.ReportDocumentStudentScore) *ReportDocumentStudentScoreOutput {
	student := user_dto.ToStudentOutput(&studentScore.Student)

	return &ReportDocumentStudentScoreOutput{
		ID:               studentScore.ID,
		StudentID:        studentScore.StudentID,
		Student:          student,
		Score:            studentScore.Score,
		ReportDocumentID: studentScore.ReportDocumentID,
		CreatedAt:        studentScore.CreatedAt,
		UpdatedAt:        studentScore.UpdatedAt,
	}
}

type ReportDocumentOutput struct {
	ID                 int64     `json:"id"`
	Name               string    `json:"name"`
	FileIDs            []string  `json:"file_ids"`
	CapstoneGroupID    int64     `json:"capstone_group_id"`
	MentorReviewStatus string    `json:"mentor_review_status"`
	TypeReport         string    `json:"type_report"`
	Conclusion         *string   `json:"conclusion"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func ToReportDocumentOutput(reportDocument *model.ReportDocument) *ReportDocumentOutput {
	return &ReportDocumentOutput{
		ID:                 reportDocument.ID,
		Name:               reportDocument.Name,
		FileIDs:            reportDocument.FileIDs,
		CapstoneGroupID:    reportDocument.CapstoneGroupID,
		MentorReviewStatus: reportDocument.MentorReviewStatus,
		TypeReport:         reportDocument.TypeReport,
		Conclusion:         reportDocument.Conclusion,
		CreatedAt:          reportDocument.CreatedAt,
		UpdatedAt:          reportDocument.UpdatedAt,
	}
}

type CapstoneGroupOutput struct {
	ID         int64     `json:"id"`
	NameGroup  string    `json:"name_group"`
	TopicID    *int64    `json:"topic_id"`
	MajorID    int64     `json:"major_id"`
	SemesterID int64     `json:"semester_id"`
	LeaderID   int64     `json:"leader_id"`
	MentorID   *int64    `json:"mentor_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CapstoneGroupWithTotalMemberOutput struct {
	ID           int64     `json:"id"`
	NameGroup    string    `json:"name_group"`
	TopicID      *int64    `json:"topic_id"`
	MajorID      int64     `json:"major_id"`
	SemesterID   int64     `json:"semester_id"`
	LeaderID     int64     `json:"leader_id"`
	MentorID     *int64    `json:"mentor_id"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	TotalMembers int64     `json:"total_members"`
}

func ToCapstoneGroupOutput(capstoneGroup *model.CapstoneGroup) CapstoneGroupOutput {
	return CapstoneGroupOutput{
		ID:         capstoneGroup.ID,
		NameGroup:  capstoneGroup.NameGroup,
		TopicID:    capstoneGroup.TopicID,
		MajorID:    capstoneGroup.MajorID,
		SemesterID: capstoneGroup.SemesterID,
		LeaderID:   capstoneGroup.LeaderID,
		MentorID:   capstoneGroup.MentorID,
		Status:     capstoneGroup.Status,
		CreatedAt:  capstoneGroup.CreatedAt,
		UpdatedAt:  capstoneGroup.UpdatedAt,
	}
}

func ToCapstoneGroupWithTotalMemberOutput(capstoneGroup *model.CapstoneGroupWithTotalMember) CapstoneGroupWithTotalMemberOutput {
	return CapstoneGroupWithTotalMemberOutput{
		ID:           capstoneGroup.ID,
		NameGroup:    capstoneGroup.NameGroup,
		TopicID:      capstoneGroup.TopicID,
		MajorID:      capstoneGroup.MajorID,
		SemesterID:   capstoneGroup.SemesterID,
		LeaderID:     capstoneGroup.LeaderID,
		MentorID:     capstoneGroup.MentorID,
		Status:       capstoneGroup.Status,
		CreatedAt:    capstoneGroup.CreatedAt,
		UpdatedAt:    capstoneGroup.UpdatedAt,
		TotalMembers: capstoneGroup.TotalMembers,
	}
}

type InvitationMentorCapstoneGroupOutput struct {
	ID              int64                   `json:"id"`
	Status          string                  `json:"status"`
	CapstoneGroupID int64                   `json:"capstone_group_id"`
	MentorID        int64                   `json:"mentor_id"`
	Mentor          *user_dto.TeacherOutput `json:"mentor"`
	ExpiredAt       time.Time               `json:"expired_at"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
}

func ToInvitationMentorCapstoneGroupOutput(invitationMentorCapstoneGroup *model.InvitationMentorCapstoneGroup) *InvitationMentorCapstoneGroupOutput {
	return &InvitationMentorCapstoneGroupOutput{
		ID:              invitationMentorCapstoneGroup.ID,
		Status:          invitationMentorCapstoneGroup.Status,
		CapstoneGroupID: invitationMentorCapstoneGroup.CapstoneGroupID,
		MentorID:        invitationMentorCapstoneGroup.MentorID,
		Mentor:          user_dto.ToTeacherOutput(&invitationMentorCapstoneGroup.Mentor),
		ExpiredAt:       invitationMentorCapstoneGroup.ExpiredAt,
		CreatedAt:       invitationMentorCapstoneGroup.CreatedAt,
		UpdatedAt:       invitationMentorCapstoneGroup.UpdatedAt,
	}
}

type CapstoneGroupReviewOutput struct {
	ID               int64     `json:"id"`
	CapstoneGroupID  int64     `json:"capstone_group_id"`
	ScheduleReviewID int64     `json:"schedule_review_id"`
	ReportFiles      []string  `json:"report_files"`
	Feedback         string    `json:"feedback"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func ToCapstoneGroupReviewOutput(capstoneGroupReview *model.CapstoneGroupReview) *CapstoneGroupReviewOutput {
	return &CapstoneGroupReviewOutput{
		ID:               capstoneGroupReview.ID,
		CapstoneGroupID:  capstoneGroupReview.CapstoneGroupID,
		ScheduleReviewID: capstoneGroupReview.ScheduleReviewID,
		ReportFiles:      capstoneGroupReview.ReportFiles,
		Feedback:         capstoneGroupReview.Feedback,
		CreatedAt:        capstoneGroupReview.CreatedAt,
		UpdatedAt:        capstoneGroupReview.UpdatedAt,
	}
}

type ListInvitationMentorCapstoneGroupOutput struct {
	Meta  dto.MetaPagination                     `json:"meta"`
	Items []*InvitationMentorCapstoneGroupOutput `json:"items"`
}

type ListCapstoneGroupOutput struct {
	Meta  dto.MetaPagination    `json:"meta"`
	Items []CapstoneGroupOutput `json:"items"`
}

type ListCapstoneGroupWithTotalMemberOutput struct {
	Meta  dto.MetaPagination                   `json:"meta"`
	Items []CapstoneGroupWithTotalMemberOutput `json:"items"`
}

type MentorAndListMemberCapstoneGroupOutput struct {
	Mentor   *user_dto.TeacherOutput   `json:"mentor"`
	Members  []*user_dto.StudentOutput `json:"members"`
	LeaderID int64                     `json:"leader_id"`
}

// This used for swagger
type GetCapstoneGroupSwaggerOutput struct {
	Code    int                  `json:"code"`
	Success bool                 `json:"message"`
	Data    *CapstoneGroupOutput `json:"data"`
}

type GetCapstoneGroupReviewSwaggerOutput struct {
	Code    int                       `json:"code"`
	Success bool                      `json:"message"`
	Data    CapstoneGroupReviewOutput `json:"data"`
}

type MentorAndListMemberCapstoneGroupSwaggerOutput struct {
	Code    int                                     `json:"code"`
	Success bool                                    `json:"message"`
	Data    *MentorAndListMemberCapstoneGroupOutput `json:"data"`
}

type ListStudentHaveCapstoneGroupSwaggerOutput struct {
	Code    int                        `json:"code"`
	Success bool                       `json:"message"`
	Data    *[]*user_dto.StudentOutput `json:"data"`
}

type ReportDocumentSwaggerOutput struct {
	Code    int                  `json:"code"`
	Success bool                 `json:"message"`
	Data    ReportDocumentOutput `json:"data"`
}

type ReportDocumentsSwaggerOutput struct {
	Code    int                    `json:"code"`
	Success bool                   `json:"message"`
	Data    []ReportDocumentOutput `json:"data"`
}

type ReportDocumentStudentsScoreSwaggerOutput struct {
	Code    int                                `json:"code"`
	Success bool                               `json:"message"`
	Data    []ReportDocumentStudentScoreOutput `json:"data"`
}

type ReportCommentWithUserInfoOutput struct {
	ID               int64                `json:"id"`
	UserID           int64                `json:"user_id"`
	User             *user_dto.UserOutput `json:"user"`
	ReportDocumentID int64                `json:"report_document_id"`
	Message          string               `json:"message"`
	GroupComment     int64                `json:"group_comment"`
	CreatedAt        time.Time            `json:"created_at"`
	UpdatedAt        time.Time            `json:"updated_at"`
}

func ToReportCommentWithUserInfoOutput(reportComment *model.ReportComment) *ReportCommentWithUserInfoOutput {
	return &ReportCommentWithUserInfoOutput{
		ID:               reportComment.ID,
		UserID:           reportComment.UserID,
		User:             user_dto.ToUserOutput(&reportComment.User),
		ReportDocumentID: reportComment.ReportDocumentID,
		Message:          reportComment.Message,
		GroupComment:     reportComment.GroupComment,
		CreatedAt:        reportComment.CreatedAt,
		UpdatedAt:        reportComment.UpdatedAt,
	}
}

type ReportCommentsSwaggerOutput struct {
	Code    int                               `json:"code"`
	Success bool                              `json:"message"`
	Data    []ReportCommentWithUserInfoOutput `json:"data"`
}

type CurrentListCapstoneGroupSwaggerOutput struct {
	Code    int                                  `json:"code"`
	Success bool                                 `json:"message"`
	Data    []CapstoneGroupWithTotalMemberOutput `json:"data"`
}
