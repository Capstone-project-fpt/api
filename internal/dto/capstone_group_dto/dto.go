package capstone_group_dto

type CreateCapstoneGroupInput struct {
	NameGroup  string  `json:"name_group" binding:"required"`
	StudentIds []int64 `json:"student_ids" binding:"required"`
	SemesterID int64   `json:"semester_id" binding:"required"`
	MajorID    int64   `json:"major_id" binding:"required"`
}

type UpdateCapstoneGroupInput struct {
	ID        int64  `json:"id" binding:"required"`
	NameGroup string `json:"name_group" binding:"required"`
}

type InviteMentorToCapstoneGroupInput struct {
	TeacherID       int64 `json:"teacher_id" binding:"required"`
	SemesterID      int64 `json:"semester_id" binding:"required"`
	CapstoneGroupID int64 `swaggerignore:"true"`
}

type ResponseInviteMentorToCapstoneGroupInput struct {
	Token           string `json:"token" binding:"required"`
	Status          string `json:"status" binding:"required" validate:"required,oneof=approve reject"`
	CapstoneGroupID int64  `swaggerignore:"true"`
}

type UpdateReportFilesCapstoneGroupReviewInput struct {
	CapstoneGroupReviewID int64    `json:"capstone_group_review_id" binding:"required"`
	ReportFiles           []string `json:"report_files" binding:"required"`
	CapstoneGroupID       int64    `swaggerignore:"true"`
}

type FeedbackCapstoneGroupReviewInput struct {
	CapstoneGroupReviewID int64  `json:"capstone_group_review_id" binding:"required"`
	Feedback              string `json:"feedback"`
	CapstoneGroupID       int64  `swaggerignore:"true"`
}

type CreateCapstoneGroupReportDocumentInput struct {
	Name            string   `json:"name" binding:"required"`
	CapstoneGroupID int64    `swaggerignore:"true"`
	FileIDs         []string `json:"file_ids" binding:"required"`
	TypeReport      string   `json:"type_report" binding:"required" validate:"required,oneof=first_report second_report third_report fourth_report fifth_report sixth_report seventh_report"`
}

type UpdateCapstoneGroupReportDocumentInput struct {
	ID              int64    `json:"id" binding:"required"`
	Name            string   `json:"name" binding:"required"`
	CapstoneGroupID int64    `swaggerignore:"true"`
	FileIDs         []string `json:"file_ids" binding:"required"`
}

type StudentScoreReportDocumentInput struct {
	Score     float64 `json:"score" binding:"required"`
	StudentID int64   `json:"student_id" binding:"required"`
}

type MentorUpdateStudentScoreForReportDocument struct {
	ReportDocumentID int64                             `json:"report_document_id" binding:"required"`
	StudentScoreData []StudentScoreReportDocumentInput `json:"student_score_data" binding:"required"`
	Conclusion       *string                           `json:"conclusion"`
	CapstoneGroupID  int64                             `swaggerignore:"true"`
}

type AdminUpdateStudentScoreForReportDocument struct {
	ID    int64   `json:"id" binding:"required"`
	Score float64 `json:"score" binding:"required"`
}

type UpdateCapstoneGroupStudentInput struct {
	ID         int64   `swaggerignore:"true"`
	StudentIDs []int64 `json:"student_ids" binding:"required"`
}

type CommentReportInput struct {
	Message          string `json:"message" binding:"required"`
	CapstoneGroupID  int64  `swaggerignore:"true"`
	ReportDocumentID int64  `swaggerignore:"true"`
}

type UpdateCommentReportInput struct {
	ID               int64  `json:"id" binding:"required"`
	Message          string `json:"message" binding:"required"`
	CapstoneGroupID  int64  `swaggerignore:"true"`
	ReportDocumentID int64  `swaggerignore:"true"`
}

type DeleteCommentReportInput struct {
	ID               int64 `json:"id" binding:"required"`
	CapstoneGroupID  int64 `swaggerignore:"true"`
	ReportDocumentID int64 `swaggerignore:"true"`
}

type GetListCommentReportInput struct {
	OrderBy          *string `form:"order_by" example:"ASC|DESC"`
	ReportDocumentID int64   `swaggerignore:"true"`
	CapstoneGroupID  int64   `swaggerignore:"true"`
}
