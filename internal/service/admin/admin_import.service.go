package service

import (
	"errors"
	"mime/multipart"
	"net/http"
	"sort"
	"time"

	"github.com/api/database/model"
	"github.com/api/global"
	"github.com/api/internal/constant"
	"github.com/api/internal/dto/import_dto"
	file_util "github.com/api/pkg/utils/file"
	password_util "github.com/api/pkg/utils/password"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/thoas/go-funk"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
)

type UserInfo struct {
	Name        string
	UserType    string
	Password    string
	Email       string
	PhoneNumber string
}

type StudentInfo struct {
	Code       string
	SubMajorID any
}

type TeacherInfo struct {
	SubMajorID any
}

type UserStudentInfo struct {
	Row     int
	User    UserInfo
	Student StudentInfo
}

type UserTeacherInfo struct {
	Row     int
	User    UserInfo
	Teacher TeacherInfo
}

func (as *adminService) UploadFileTeacherData(ctx *gin.Context, fileUpload *multipart.FileHeader) (int, *import_dto.ImportOutput) {
	resultImport := import_dto.ImportOutput{
		SuccessCount:     0,
		FailedCount:      0,
		FailedImportDocs: []import_dto.FailedImportRecordOutput{},
		Exception:        "",
	}

	rows, err := as.openAndGetAllRowOfExcelFile(ctx, fileUpload)
	if err != nil {
		resultImport.Exception = err.Error()

		return http.StatusBadRequest, &resultImport
	}

	if len(rows) == 1 {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.NotAllowEmptyDataInFile,
		})
		resultImport.Exception = message

		return http.StatusBadRequest, &resultImport
	}

	redis := global.RDb
	_, err = redis.Get(ctx, constant.LockProcessType.CreateTeacherAccount).Result()
	if err == nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.OtherSessionImportTeacherInProcess,
		})
		resultImport.Exception = message

		return http.StatusBadRequest, &resultImport
	}

	var subMajors []model.SubMajor
	err = global.Db.Model(&model.SubMajor{}).Select("id").Find(&subMajors).Error
	if err != nil {
		global.Logger.Error("Failed to get all sub majors: ", zap.Error(err))
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		resultImport.Exception = message

		return http.StatusBadRequest, &resultImport
	}

	var subMajorIds []int
	for _, subMajor := range subMajors {
		subMajorIds = append(subMajorIds, int(subMajor.ID))
	}

	teacherDataImportMapping := constant.TeacherDataImportMapping
	checkImport := new(file_util.CheckValidImport)
	var userTeacherInfos []UserTeacherInfo

	_, err = redis.Set(ctx, constant.LockProcessType.CreateStudentAccount, true, time.Duration(300)*time.Second).Result() // 5 minutes
	if err != nil {
		global.Logger.Error("Failed to set lock process type: ", zap.Error(err))
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		resultImport.Exception = message

		return http.StatusBadRequest, &resultImport
	}

	for rowIndex := 1; rowIndex < len(rows); rowIndex++ {
		row := rows[rowIndex]

		if len(row) < constant.TotalColumnTeacherImportData {
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.InvalidFile,
			})
			resultImport.Exception = message

			return http.StatusBadRequest, &resultImport
		}

		name := row[teacherDataImportMapping.Name]
		_, err := checkImport.CheckImportData(file_util.CheckImportDataInput{
			ColumnName:         "Name",
			CellData:           name,
			RowNum:             rowIndex + 1,
			IsRequired:         true,
			CustomMessageError: "",
			ExpectedType:       "string",
		})
		if err != nil {
			resultImport.FailedCount++
			resultImport.FailedImportDocs = append(resultImport.FailedImportDocs, import_dto.FailedImportRecordOutput{
				Row:   rowIndex + 1,
				Error: err.Error(),
			})
		}

		email := row[teacherDataImportMapping.Email]
		_, err = checkImport.CheckImportData(file_util.CheckImportDataInput{
			ColumnName:         "Email",
			CellData:           email,
			RowNum:             rowIndex + 1,
			IsRequired:         true,
			CustomMessageError: "",
			ExpectedType:       "string",
		})
		if err != nil {
			resultImport.FailedCount++
			resultImport.FailedImportDocs = append(resultImport.FailedImportDocs, import_dto.FailedImportRecordOutput{
				Row:   rowIndex + 1,
				Error: err.Error(),
			})
		}

		phoneNumber := row[teacherDataImportMapping.PhoneNumber]
		_, err = checkImport.CheckImportData(file_util.CheckImportDataInput{
			ColumnName:         "PhoneNumber",
			CellData:           phoneNumber,
			RowNum:             rowIndex + 1,
			IsRequired:         true,
			CustomMessageError: "",
			ExpectedType:       "string",
		})
		if err != nil {
			resultImport.FailedCount++
			resultImport.FailedImportDocs = append(resultImport.FailedImportDocs, import_dto.FailedImportRecordOutput{
				Row:   rowIndex + 1,
				Error: err.Error(),
			})
		}
		subMajorIDData := row[teacherDataImportMapping.SubMajorID]
		subMajorID, err := checkImport.CheckImportData(file_util.CheckImportDataInput{
			ColumnName:         "SubMajorID",
			CellData:           subMajorIDData,
			RowNum:             rowIndex + 1,
			IsRequired:         true,
			CustomMessageError: "",
			ExpectedType:       "number",
		})
		if err != nil {
			resultImport.FailedCount++
			resultImport.FailedImportDocs = append(resultImport.FailedImportDocs, import_dto.FailedImportRecordOutput{
				Row:   rowIndex + 1,
				Error: err.Error(),
			})
		}
		if err == nil && !funk.Contains(subMajorIds, subMajorID.(int)) {
			resultImport.FailedCount++
			resultImport.FailedImportDocs = append(resultImport.FailedImportDocs, import_dto.FailedImportRecordOutput{
				Row:   rowIndex + 1,
				Error: checkImport.GetInvalidMessage("SubMajorID", rowIndex+1),
			})
		}

		password := password_util.GenerateRandomPassword(constant.DefaultPasswordLength)

		userTeacherInfos = append(userTeacherInfos, UserTeacherInfo{
			Row: rowIndex + 1,
			User: UserInfo{
				Name:        name,
				UserType:    constant.UserType.Teacher,
				Password:    password,
				Email:       email,
				PhoneNumber: phoneNumber,
			},
			Teacher: TeacherInfo{
				SubMajorID: subMajorID,
			},
		})
	}

	var teacherImportEmails []string
	for _, userTeacherInfo := range userTeacherInfos {
		teacherImportEmails = append(teacherImportEmails, userTeacherInfo.User.Email)
	}

	var teachersExit []model.User
	if err := global.Db.Model(&model.User{}).Where("email IN ?", teacherImportEmails).Select("email").Find(&teachersExit).Error; err != nil {
		return http.StatusInternalServerError, nil
	}

	var teacherEmailsExit []string
	for _, teacher := range teachersExit {
		teacherEmailsExit = append(teacherEmailsExit, teacher.Email)
	}
	for _, userTeacherInfo := range userTeacherInfos {
		if funk.Contains(teacherEmailsExit, userTeacherInfo.User.Email) {
			resultImport.FailedCount++
			resultImport.FailedImportDocs = append(resultImport.FailedImportDocs, import_dto.FailedImportRecordOutput{
				Row:   userTeacherInfo.Row,
				Error: "Email giáo viên đã tồn tại",
			})
		}
	}
	statusCode := http.StatusBadRequest
	resultImportAfterProcess := &resultImport

	if resultImport.FailedCount == 0 {
		statusCode, resultImportAfterProcess = as.transactionCreateTeacherAccount(ctx, userTeacherInfos, &resultImport)
	}

	sort.Slice(resultImportAfterProcess.FailedImportDocs, func(i, j int) bool {
		return resultImportAfterProcess.FailedImportDocs[i].Row < resultImportAfterProcess.FailedImportDocs[j].Row
	})

	// Close Lock
	err = redis.Del(ctx, constant.LockProcessType.CreateStudentAccount).Err()
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		resultImport.Exception = message
		statusCode = http.StatusInternalServerError
	}

	return statusCode, resultImportAfterProcess
}

func (as *adminService) UploadFileStudentData(ctx *gin.Context, fileUpload *multipart.FileHeader) (int, *import_dto.ImportOutput) {
	resultImport := import_dto.ImportOutput{
		SuccessCount:     0,
		FailedCount:      0,
		FailedImportDocs: []import_dto.FailedImportRecordOutput{},
		Exception:        "",
	}

	rows, err := as.openAndGetAllRowOfExcelFile(ctx, fileUpload)
	if err != nil {
		resultImport.Exception = err.Error()

		return http.StatusBadRequest, &resultImport
	}

	if len(rows) == 1 {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.NotAllowEmptyDataInFile,
		})
		resultImport.Exception = message

		return http.StatusBadRequest, &resultImport
	}

	redis := global.RDb
	_, err = redis.Get(ctx, constant.LockProcessType.CreateStudentAccount).Result()
	if err == nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.OtherSessionImportStudentInProcess,
		})
		resultImport.Exception = message

		return http.StatusBadRequest, &resultImport
	}

	var subMajors []model.SubMajor
	err = global.Db.Model(&model.SubMajor{}).Select("id").Find(&subMajors).Error
	if err != nil {
		global.Logger.Error("Failed to get all sub majors: ", zap.Error(err))
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		resultImport.Exception = message

		return http.StatusBadRequest, &resultImport
	}

	var subMajorIds []int
	for _, subMajor := range subMajors {
		subMajorIds = append(subMajorIds, int(subMajor.ID))
	}

	studentDataImportMapping := constant.StudentDataImportMapping
	checkImport := new(file_util.CheckValidImport)
	var userStudentInfos []UserStudentInfo

	_, err = redis.Set(ctx, constant.LockProcessType.CreateStudentAccount, true, time.Duration(300)*time.Second).Result() // 5 minutes
	if err != nil {
		global.Logger.Error("Failed to set lock process type: ", zap.Error(err))
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		resultImport.Exception = message

		return http.StatusBadRequest, &resultImport
	}

	for rowIndex := 1; rowIndex < len(rows); rowIndex++ {
		row := rows[rowIndex]

		if len(row) < constant.TotalColumnStudentImportData {
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.InvalidFile,
			})
			resultImport.Exception = message

			return http.StatusBadRequest, &resultImport
		}

		name := row[studentDataImportMapping.Name]
		_, err := checkImport.CheckImportData(file_util.CheckImportDataInput{
			ColumnName:         "Name",
			CellData:           name,
			RowNum:             rowIndex + 1,
			IsRequired:         true,
			CustomMessageError: "",
			ExpectedType:       "string",
		})
		if err != nil {
			resultImport.FailedCount++
			resultImport.FailedImportDocs = append(resultImport.FailedImportDocs, import_dto.FailedImportRecordOutput{
				Row:   rowIndex + 1,
				Error: err.Error(),
			})
		}

		email := row[studentDataImportMapping.Email]
		_, err = checkImport.CheckImportData(file_util.CheckImportDataInput{
			ColumnName:         "Email",
			CellData:           email,
			RowNum:             rowIndex + 1,
			IsRequired:         true,
			CustomMessageError: "",
			ExpectedType:       "string",
		})
		if err != nil {
			resultImport.FailedCount++
			resultImport.FailedImportDocs = append(resultImport.FailedImportDocs, import_dto.FailedImportRecordOutput{
				Row:   rowIndex + 1,
				Error: err.Error(),
			})
		}

		code := row[studentDataImportMapping.Code]
		_, err = checkImport.CheckImportData(file_util.CheckImportDataInput{
			ColumnName:         "Code",
			CellData:           code,
			RowNum:             rowIndex + 1,
			IsRequired:         true,
			CustomMessageError: "",
			ExpectedType:       "string",
		})
		if err != nil {
			resultImport.FailedCount++
			resultImport.FailedImportDocs = append(resultImport.FailedImportDocs, import_dto.FailedImportRecordOutput{
				Row:   rowIndex + 1,
				Error: err.Error(),
			})
		}

		phoneNumber := row[studentDataImportMapping.PhoneNumber]
		_, err = checkImport.CheckImportData(file_util.CheckImportDataInput{
			ColumnName:         "PhoneNumber",
			CellData:           phoneNumber,
			RowNum:             rowIndex + 1,
			IsRequired:         true,
			CustomMessageError: "",
			ExpectedType:       "string",
		})
		if err != nil {
			resultImport.FailedCount++
			resultImport.FailedImportDocs = append(resultImport.FailedImportDocs, import_dto.FailedImportRecordOutput{
				Row:   rowIndex + 1,
				Error: err.Error(),
			})
		}
		subMajorIDData := row[studentDataImportMapping.SubMajorID]
		subMajorID, err := checkImport.CheckImportData(file_util.CheckImportDataInput{
			ColumnName:         "SubMajorID",
			CellData:           subMajorIDData,
			RowNum:             rowIndex + 1,
			IsRequired:         true,
			CustomMessageError: "",
			ExpectedType:       "number",
		})
		if err != nil {
			resultImport.FailedCount++
			resultImport.FailedImportDocs = append(resultImport.FailedImportDocs, import_dto.FailedImportRecordOutput{
				Row:   rowIndex + 1,
				Error: err.Error(),
			})
		}
		if err == nil && !funk.Contains(subMajorIds, subMajorID.(int)) {
			resultImport.FailedCount++
			resultImport.FailedImportDocs = append(resultImport.FailedImportDocs, import_dto.FailedImportRecordOutput{
				Row:   rowIndex + 1,
				Error: checkImport.GetInvalidMessage("SubMajorID", rowIndex+1),
			})
		}

		password := password_util.GenerateRandomPassword(constant.DefaultPasswordLength)

		userStudentInfos = append(userStudentInfos, UserStudentInfo{
			Row: rowIndex + 1,
			User: UserInfo{
				Name:        name,
				UserType:    constant.UserType.Student,
				Password:    password,
				Email:       email,
				PhoneNumber: phoneNumber,
			},
			Student: StudentInfo{
				Code:       code,
				SubMajorID: subMajorID,
			},
		})
	}

	var studentImportCodes []string
	for _, userStudentInfo := range userStudentInfos {
		studentImportCodes = append(studentImportCodes, userStudentInfo.Student.Code)
	}

	var studentsExist []model.Student
	if err := global.Db.Model(&model.Student{}).Where("code IN ?", studentImportCodes).Select("code").Find(&studentsExist).Error; err != nil {
		return http.StatusInternalServerError, nil
	}
	var studentCodesExist []string
	for _, student := range studentsExist {
		studentCodesExist = append(studentCodesExist, student.Code)
	}
	for _, userStudentInfo := range userStudentInfos {
		if funk.Contains(studentCodesExist, userStudentInfo.Student.Code) {
			resultImport.FailedCount++
			resultImport.FailedImportDocs = append(resultImport.FailedImportDocs, import_dto.FailedImportRecordOutput{
				Row:   userStudentInfo.Row,
				Error: "Mã sinh viên đã tồn tại",
			})
		}
	}

	statusCode := http.StatusBadRequest
	resultImportAfterProcess := &resultImport

	if resultImport.FailedCount == 0 {
		statusCode, resultImportAfterProcess = as.transactionCreateStudentAccount(ctx, userStudentInfos, &resultImport)
	}

	sort.Slice(resultImportAfterProcess.FailedImportDocs, func(i, j int) bool {
		return resultImportAfterProcess.FailedImportDocs[i].Row < resultImportAfterProcess.FailedImportDocs[j].Row
	})

	// Close Lock
	err = redis.Del(ctx, constant.LockProcessType.CreateStudentAccount).Err()
	if err != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		resultImport.Exception = message
		statusCode = http.StatusInternalServerError
	}

	return statusCode, resultImportAfterProcess
}

func (as *adminService) openAndGetAllRowOfExcelFile(ctx *gin.Context, fileUpload *multipart.FileHeader) ([][]string, error) {
	file, err := fileUpload.Open()
	if err != nil {
		global.Logger.Error("Failed to read file: ", zap.Error(err))
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidFile,
		})

		return nil, errors.New(message)
	}
	defer file.Close()

	fileContent, err := excelize.OpenReader(file)

	if err != nil {
		global.Logger.Error("Failed to read file: ", zap.Error(err))
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidFile,
		})

		return nil, errors.New(message)
	}
	defer fileContent.Close()

	rows, err := fileContent.GetRows("Sheet1")
	if err != nil {
		global.Logger.Error("Failed to get all rows: ", zap.Error(err))
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InvalidFile,
		})

		return nil, errors.New(message)
	}

	return rows, nil
}

func (as *adminService) transactionCreateTeacherAccount(ctx *gin.Context, userTeacherInfos []UserTeacherInfo, resultImport *import_dto.ImportOutput) (int, *import_dto.ImportOutput) {
	tx := global.Db.Begin()
	if tx.Error != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		global.Logger.Error("Failed to begin a transaction, Error: ", zap.Error(tx.Error))

		resultImport.Exception = message
		return http.StatusInternalServerError, resultImport
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.InternalServerError,
			})
			global.Logger.Error("Failed to recover a transaction, Error: ", zap.Error(tx.Error))

			resultImport.Exception = message
		}
	}()

	var role model.Role
	if err := tx.Model(&model.Role{}).Select("id").First(&role, "name = ?", constant.RoleType.Teacher).Error; err != nil {
		tx.Rollback()
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		resultImport.Exception = message
		return http.StatusInternalServerError, resultImport
	}

	for _, userTeacherInfo := range userTeacherInfos {
		hashPassword, err := password_util.HashPassword(userTeacherInfo.User.Password)
		if err != nil {
			tx.Rollback()
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.InternalServerError,
			})
			resultImport.Exception = message
			return http.StatusInternalServerError, resultImport
		}

		user := model.User{
			Name:        userTeacherInfo.User.Name,
			UserType:    userTeacherInfo.User.UserType,
			Email:       userTeacherInfo.User.Email,
			Password:    hashPassword,
			PhoneNumber: userTeacherInfo.User.PhoneNumber,
		}
		if err := tx.Save(&user).Error; err != nil {
			tx.Rollback()
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.InternalServerError,
			})
			resultImport.Exception = message
			return http.StatusInternalServerError, resultImport
		}

		teacher := model.Teacher{
			SubMajorID: int64(userTeacherInfo.Teacher.SubMajorID.(int)),
			UserID:     user.ID,
		}

		if err := tx.Save(&teacher).Error; err != nil {
			tx.Rollback()
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.InternalServerError,
			})
			resultImport.Exception = message
			return http.StatusInternalServerError, resultImport
		}

		if err := tx.Exec("INSERT INTO users_roles (user_id, role_id) VALUES (?, ?)", user.ID, role.ID).Error; err != nil {
			tx.Rollback()
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.InternalServerError,
			})
			resultImport.Exception = message
			return http.StatusInternalServerError, resultImport
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		resultImport.Exception = message
		return http.StatusInternalServerError, resultImport
	}

	// TODO: Handle send email

	return http.StatusOK, resultImport
}

func (as *adminService) transactionCreateStudentAccount(ctx *gin.Context, userStudentInfos []UserStudentInfo, resultImport *import_dto.ImportOutput) (int, *import_dto.ImportOutput) {
	tx := global.Db.Begin()
	if tx.Error != nil {
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		global.Logger.Error("Failed to begin a transaction, Error: ", zap.Error(tx.Error))

		resultImport.Exception = message
		return http.StatusInternalServerError, resultImport
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.InternalServerError,
			})
			global.Logger.Error("Failed to recover a transaction, Error: ", zap.Error(tx.Error))

			resultImport.Exception = message
		}
	}()

	var role model.Role
	if err := tx.Model(&model.Role{}).Select("id").First(&role, "name = ?", constant.RoleType.Student).Error; err != nil {
		tx.Rollback()
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		resultImport.Exception = message
		return http.StatusInternalServerError, resultImport
	}

	for _, userStudentInfo := range userStudentInfos {
		hashPassword, err := password_util.HashPassword(userStudentInfo.User.Password)
		if err != nil {
			tx.Rollback()
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.InternalServerError,
			})
			resultImport.Exception = message
			return http.StatusInternalServerError, resultImport
		}

		user := model.User{
			Name:        userStudentInfo.User.Name,
			UserType:    userStudentInfo.User.UserType,
			Email:       userStudentInfo.User.Email,
			Password:    hashPassword,
			PhoneNumber: userStudentInfo.User.PhoneNumber,
		}
		if err := tx.Save(&user).Error; err != nil {
			tx.Rollback()
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.InternalServerError,
			})
			resultImport.Exception = message
			return http.StatusInternalServerError, resultImport
		}

		student := model.Student{
			Code:       userStudentInfo.Student.Code,
			SubMajorID: int64(userStudentInfo.Student.SubMajorID.(int)),
			UserID:     user.ID,
		}

		if err := tx.Save(&student).Error; err != nil {
			tx.Rollback()
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.InternalServerError,
			})
			resultImport.Exception = message
			return http.StatusInternalServerError, resultImport
		}

		if err := tx.Exec("INSERT INTO users_roles (user_id, role_id) VALUES (?, ?)", user.ID, role.ID).Error; err != nil {
			tx.Rollback()
			message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: constant.MessageI18nId.InternalServerError,
			})
			resultImport.Exception = message
			return http.StatusInternalServerError, resultImport
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		message := global.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: constant.MessageI18nId.InternalServerError,
		})
		resultImport.Exception = message
		return http.StatusInternalServerError, resultImport
	}

	// TODO: Handle send email

	return http.StatusOK, resultImport
}
