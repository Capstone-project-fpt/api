package upload_dto

type GenerateUploadPresignUrlInput struct {
	Key string
}

type GenerateUploadPresignUrlsInput struct {
	Key []string `json:"key"`
}

// This used for swagger
type GenerateUploadPresignUrlsSwaggerOutput struct {
	Code    int      `json:"code"`
	Success bool     `json:"message"`
	Data    []string `json:"data"`
}
