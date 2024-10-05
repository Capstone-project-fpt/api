package service

import (
	"errors"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/user_dto"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type IUserService interface {
	GetUser(ctx *gin.Context, userID int) (*user_dto.GetUserOutput, error)
}

type userService struct {
}

func NewUserService() IUserService {
	return &userService{}
}

func (u *userService) GetUser(ctx *gin.Context, userID int) (*user_dto.GetUserOutput, error) {
	messageUserNotfound := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: constant.MessageI18nId.UserNotFound,
	})
	var user model.User
	if err := global.Db.Model(&model.User{}).Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, errors.New(messageUserNotfound)
	}

	var commonUserInfo user_dto.UserOutput = user_dto.UserOutput{
		ID:          int(user.ID),
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		UserType:    user.UserType,
	}
	var extraInfo user_dto.ExtraInfo

	switch user.UserType {
	case constant.UserType.Student:
		var student model.Student
		if err := global.Db.Model(&model.Student{}).Omit("UserID", "UpdatedAt").Where("user_id = ?", user.ID).First(&student).Error; err != nil {
			return nil, errors.New(messageUserNotfound)
		}

		extraInfo.Student = &user_dto.StudentInfoOutput{
			StudentID:       int(student.ID),
			Code:            student.Code,
			SubMajorId:      int(student.SubMajorID),
			CapstoneGroupID: int(student.CapstoneGroupID),
			CreatedAt:       student.CreatedAt,
		}
	case constant.UserType.Teacher:
		var teacher model.Teacher
		if err := global.Db.Model(&model.Teacher{}).Omit("UserID", "UpdatedAt").Where("user_id = ?", user.ID).First(&teacher).Error; err != nil {
			return nil, errors.New(messageUserNotfound)
		}

		extraInfo.Teacher = &user_dto.TeacherInfoOutput{
			TeacherID:  int(teacher.ID),
			SubMajorID: int(teacher.SubMajorID),
			CreatedAt:  teacher.CreatedAt,
		}
	}

	return &user_dto.GetUserOutput{
		CommonInfo: &commonUserInfo,
		ExtraInfo:  &extraInfo,
	}, nil
}