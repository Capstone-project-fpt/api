package constant

const (
	Localizer                                 = "localizer"
	TotalColumnStudentImportData              = 5
	TotalColumnTeacherImportData              = 4
	DefaultPasswordLength               int64 = 8
	DefaultResetPasswordTokenLength     int64 = 8
	DefaultResetPasswordTokenExpiration int64 = 3600 // 1 hour
	DefaultInviteMentorTokenExpiration  int64 = 3600 // 1 hour
	DefaultInviteMentorTokenLength      int64 = 8
	MinTotalMemberInGroup                     = 4
	MaxTotalMemberInGroup                     = 5
	MaxTotalCapstoneGroupTeacherMentor        = 2
)

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
	ManageAccount        string
	ViewAccount          string
	ManageTopicReference string
}

var PermissionType permissionTypeType = permissionTypeType{
	ManageAccount:        "ManageAccount",
	ViewAccount:          "ViewAccount",
	ManageTopicReference: "ManageTopicReference",
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
	SendEmailCreateAccounts              string
	SendEmailInviteMentorToCapstoneGroup string
}

var SystemQueueTask systemQueueTaskType = systemQueueTaskType{
	SendEmailCreateAccounts:              "SystemTask:SendEmailCreateAccounts",
	SendEmailInviteMentorToCapstoneGroup: "SystemTask:SendEmailInviteMentorToCapstoneGroup",
}

type topicStatusReviewType struct {
	Reviewing string
	Approved  string
	Rejected  string
}

var TopicStatusReview topicStatusReviewType = topicStatusReviewType{
	Reviewing: "reviewing",
	Approved:  "approved",
	Rejected:  "rejected",
}

type capstoneGroupStatusType struct {
	ReviewingTopic string
	InProgress     string
}

var CapstoneGroupStatus capstoneGroupStatusType = capstoneGroupStatusType{
	ReviewingTopic: "reviewing_topic",
	InProgress:     "in_progress",
}

type invitationMentorCapstoneGroupType struct {
	Pending string
	Approve string
	Reject  string
}

var InvitationMentorCapstoneGroup invitationMentorCapstoneGroupType = invitationMentorCapstoneGroupType{
	Pending: "pending",
	Approve: "approve",
	Reject:  "reject",
}

type scheduleReviewType struct {
	FirstReview  string
	SecondReview string
	ThirdReview  string
}

var ScheduleReview scheduleReviewType = scheduleReviewType{
	FirstReview:  "first_review",
	SecondReview: "second_review",
	ThirdReview:  "third_review",
}

type reportDocumentType struct {
	FIRST_REPORT   string
	SECOND_REPORT  string
	THIRD_REPORT   string
	FOURTH_REPORT  string
	FIFTH_REPORT   string
	SIXTH_REPORT   string
	SEVENTH_REPORT string
}

var ReportDocument reportDocumentType = reportDocumentType{
	FIRST_REPORT:   "first_report",
	SECOND_REPORT:  "second_report",
	THIRD_REPORT:   "third_report",
	FOURTH_REPORT:  "fourth_report",
	FIFTH_REPORT:   "fifth_report",
	SIXTH_REPORT:   "sixth_report",
	SEVENTH_REPORT: "seventh_report",
}

type mentorReviewStatusReportType struct {
	Reviewing string
	Done      string
}

var MentorReviewStatusReport mentorReviewStatusReportType = mentorReviewStatusReportType{
	Reviewing: "reviewing",
	Done:      "done",
}
