package import_dto

type FailedImportRecordOutput struct {
	Row   int    `json:"row"`
	Error string `json:"error"`
}

type ImportOutput struct {
	SuccessCount     int                        `json:"success_count"`
	FailedCount      int                        `json:"failed_count"`
	FailedImportDocs []FailedImportRecordOutput `json:"failed_import_docs"`
	Exception        string                     `json:"exception"`
}
