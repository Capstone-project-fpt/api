package service

import (
	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto"
	"github.com/api/internal/dto/admin_dto"
	"github.com/api/internal/dto/user_dto"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

func (as *adminService) GetListUsers(ctx *gin.Context, input *admin_dto.GetListUsersInput) (interface{}, error) {
	var total int64
	var items []model.UserWithDetails
	getTotalQuery := global.Db.Model(&model.User{}).
		Joins("LEFT JOIN teachers ON teachers.user_id = users.id").
		Joins("LEFT JOIN students ON students.user_id = users.id")

	getUsersQuery := global.Db.Model(&model.User{}).
		Select("users.id as user_id, users.name as user_name, users.email as user_email, users.phone_number as user_phone_number, users.user_type as user_type, " +
			"COALESCE(teachers.id, 0) as teacher_id, COALESCE(teachers.sub_major_id, 0) as teacher_sub_major_id, teachers.created_at as teacher_created_at, " +
			"COALESCE(students.id, 0) as student_id, students.code as student_code, COALESCE(students.capstone_group_id, 0) as student_capstone_group_id, COALESCE(students.sub_major_id, 0) as student_sub_major_id, students.created_at as student_created_at").
		Joins("LEFT JOIN teachers ON teachers.user_id = users.id").
		Joins("LEFT JOIN students ON students.user_id = users.id")

	var userTypes []string
	userTypes = funk.Filter(input.UserTypes, func(userType string) bool {
		return userType != ""
	}).([]string)
	if len(userTypes) > 0 {
		getTotalQuery.Where("user_type IN (?)", userTypes)
		getUsersQuery.Where("user_type IN (?)", userTypes)
	}

	if err := getTotalQuery.Count(&total).Error; err != nil {
		return nil, err
	}

	if err := getUsersQuery.
		Limit(int(input.Limit)).
		Offset(int(input.Offset)).
		Scan(&items).Error; err != nil {
		return nil, err
	}

	var itemsOutput []user_dto.GetUserOutput

	for _, item := range items {
		var userOutput user_dto.GetUserOutput
		var userExtraInfo user_dto.ExtraInfo
		if item.TeacherID != 0 && item.UserType == constant.UserType.Teacher {
			teacherInfo := user_dto.TeacherInfoOutput{
				TeacherID:  int(item.TeacherID),
				SubMajorID: int(item.TeacherSubMajorID),
				CreatedAt:  item.TeacherCreatedAt,
			}
			userExtraInfo.Teacher = &teacherInfo
		}
		if item.StudentID != 0 && item.UserType == constant.UserType.Student {
			studentInfo := user_dto.StudentInfoOutput{
				StudentID:       int(item.StudentID),
				Code:            item.StudentCode,
				CapstoneGroupID: int(item.StudentCapstoneGroupID),
				SubMajorId:      int(item.StudentSubMajorID),
				CreatedAt:       item.StudentCreatedAt,
			}
			userExtraInfo.Student = &studentInfo
		}

		userOutput.ExtraInfo = &userExtraInfo

		commonInfo := user_dto.UserOutput{
			ID:          int(item.UserID),
			Name:        item.UserName,
			Email:       item.UserEmail,
			PhoneNumber: item.UserPhoneNumber,
			UserType:    item.UserType,
		}

		userOutput.CommonInfo = &commonInfo

		itemsOutput = append(itemsOutput, userOutput)
	}

	return admin_dto.ListUsersOutput{
		Meta: dto.MetaPagination{
			Total:       int(total),
			CurrentPage: int(input.Page),
		},
		Items: itemsOutput,
	}, nil
}
