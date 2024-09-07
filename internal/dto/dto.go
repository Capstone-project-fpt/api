package dto

type OutputCommon struct {
	Message string `json:"message"`
}

type MetaPagination struct {
	CurrentPage int `json:"current_page"`
	Total       int `json:"total"`
}

type OutputPagination struct {
	Meta MetaPagination `json:"meta"`
	Data interface{}    `json:"data"`
}
