package topic_reference_dto

type TeacherCreateTopicReferenceInput struct {
	Name string `json:"name" binding:"required"`
	Path string `json:"path" binding:"required"`
}

type TeacherUpdateTopicReferenceInput struct {
	ID   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Path string `json:"path" binding:"required"`
}

type TeacherDeleteTopicReferenceInput struct {
	ID int `form:"id" binding:"required" example:"1"`
}

type AdminCreateTopicReferenceInput struct {
	Name      string `json:"name" binding:"required"`
	Path      string `json:"path" binding:"required"`
	TeacherID int64  `json:"teacher_id" binding:"required"`
}
