package constant

const (
	Localizer = "localizer"
)

type MessageI18n struct {
	EmailNotFound               string
	UserNotFound                string
	TokenInvalid                string
	InternalServerError         string
	InvalidParams               string
	UserAlreadyExists           string
	InvalidStudentEmailFPT      string
	CreateStudentAccountSuccess string
	CreateTeacherAccountSuccess string
	PermissionDenied            string
	MajorNotFound               string
	SubMajorNotFound            string
}

var MessageI18nId MessageI18n = MessageI18n{
	EmailNotFound:               "EmailNotFound",
	UserNotFound:                "UserNotFound",
	TokenInvalid:                "TokenInvalid",
	InternalServerError:         "InternalServerError",
	InvalidParams:               "InvalidParams",
	UserAlreadyExists:           "UserAlreadyExists",
	InvalidStudentEmailFPT:      "InvalidStudentEmailFPT",
	CreateStudentAccountSuccess: "CreateStudentAccountSuccess",
	CreateTeacherAccountSuccess: "CreateTeacherAccountSuccess",
	PermissionDenied:            "PermissionDenied",
	MajorNotFound:               "MajorNotFound",
	SubMajorNotFound:            "SubMajorNotFound",
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

var DefaultPasswordLength int64 = 8

type RoleTypeType struct {
	Admin   string
	Student string
	Teacher string
}

var RoleType RoleTypeType = RoleTypeType{
	Admin:   "admin",
	Student: "student",
	Teacher: "teacher",
}

type PermissionTypeType struct {
	ManageAccount string
	ViewAccount   string
}

var PermissionType PermissionTypeType = PermissionTypeType{
	ManageAccount: "ManageAccount",
	ViewAccount:   "ViewAccount",
}