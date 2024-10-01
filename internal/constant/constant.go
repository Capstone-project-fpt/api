package constant

const (
	Localizer                                 = "localizer"
	TotalColumnStudentImportData              = 5
	TotalColumnTeacherImportData              = 4
	DefaultPasswordLength               int64 = 8
	DefaultResetPasswordTokenLength     int64 = 8
	DefaultResetPasswordTokenExpiration int64 = 3600 // 1 hour
)

type MessageI18n struct {
	EmailNotFound                            string
	UserNotFound                             string
	TokenInvalid                             string
	InternalServerError                      string
	InvalidParams                            string
	UserAlreadyExists                        string
	InvalidStudentEmailFPT                   string
	CreateStudentAccountSuccess              string
	CreateTeacherAccountSuccess              string
	PermissionDenied                         string
	MajorNotFound                            string
	SubMajorNotFound                         string
	AlreadySendResetPasswordLink             string
	ImportAndCreateListStudentAccountSuccess string
	InvalidFile                              string
	NotAllowEmptyDataInFile                  string
	OtherSessionImportStudentInProcess       string
	OtherSessionImportTeacherInProcess       string
}

var MessageI18nId MessageI18n = MessageI18n{
	EmailNotFound:                            "EmailNotFound",
	UserNotFound:                             "UserNotFound",
	TokenInvalid:                             "TokenInvalid",
	InternalServerError:                      "InternalServerError",
	InvalidParams:                            "InvalidParams",
	UserAlreadyExists:                        "UserAlreadyExists",
	InvalidStudentEmailFPT:                   "InvalidStudentEmailFPT",
	CreateStudentAccountSuccess:              "CreateStudentAccountSuccess",
	CreateTeacherAccountSuccess:              "CreateTeacherAccountSuccess",
	PermissionDenied:                         "PermissionDenied",
	MajorNotFound:                            "MajorNotFound",
	SubMajorNotFound:                         "SubMajorNotFound",
	AlreadySendResetPasswordLink:             "AlreadySendResetPasswordLink",
	ImportAndCreateListStudentAccountSuccess: "ImportAndCreateListStudentAccountSuccess",
	InvalidFile:                              "InvalidFile",
	NotAllowEmptyDataInFile:                  "NotAllowEmptyDataInFile",
	OtherSessionImportStudentInProcess:       "OtherSessionImportStudentInProcess",
	OtherSessionImportTeacherInProcess:       "OtherSessionImportTeacherInProcess",
}

type RedisKeyType struct {
	ActiveAccessToken  string
	ActiveRefreshToken string
}

var RedisKey RedisKeyType = RedisKeyType{
	ActiveAccessToken:  "ActiveAccessToken",
	ActiveRefreshToken: "ActiveRefreshToken",
}

type UserTypeType struct {
	Admin   string
	Student string
	Teacher string
}

var UserType UserTypeType = UserTypeType{
	Admin:   "admin",
	Student: "student",
	Teacher: "teacher",
}

type roleTypeType struct {
	Admin   string
	Student string
	Teacher string
}

var RoleType roleTypeType = roleTypeType{
	Admin:   "admin",
	Student: "student",
	Teacher: "teacher",
}

type permissionTypeType struct {
	ManageAccount string
	ViewAccount   string
}

var PermissionType permissionTypeType = permissionTypeType{
	ManageAccount: "ManageAccount",
	ViewAccount:   "ViewAccount",
}

type lockProcessTypeType struct {
	CreateStudentAccount string
	CreateTeacherAccount string
}

var LockProcessType lockProcessTypeType = lockProcessTypeType{
	CreateStudentAccount: "LockProcessCreateStudentAccount",
	CreateTeacherAccount: "LockProcessCreateTeacherAccount",
}

type StudentDataImportMappingType struct {
	Name        int
	Email       int
	Code        int
	PhoneNumber int
	SubMajorID  int
}

type TeacherDataImportMappingType struct {
	Name        int
	Email       int
	PhoneNumber int
	SubMajorID  int
}

var StudentDataImportMapping StudentDataImportMappingType = StudentDataImportMappingType{
	Name:        0,
	Email:       1,
	Code:        2,
	PhoneNumber: 3,
	SubMajorID:  4,
}

var TeacherDataImportMapping TeacherDataImportMappingType = TeacherDataImportMappingType{
	Name:        0,
	Email:       1,
	PhoneNumber: 2,
	SubMajorID:  3,
}

type systemQueueTaskType struct {
	SendEmailCreateAccounts string
}

var SystemQueueTask systemQueueTaskType = systemQueueTaskType{
	SendEmailCreateAccounts: "SystemTask:SendEmailCreateAccounts",
}
