package file_util

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"strings"
)

func IsExcelFile(file *multipart.FileHeader) bool {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".xlsx" {
		return false
	}

	return true
}

type CheckValidImport struct{}

type CheckImportDataInput struct {
	ColumnName         string
	CellData           string // Since data from Excel is always a string
	RowNum             int
	IsRequired         bool
	CustomMessageError string
	ExpectedType       string // Define the expected type (e.g., "string", "number")
}

func (c *CheckValidImport) CheckImportData(input CheckImportDataInput) (interface{}, error) {
	if input.IsRequired && input.CellData == "" {
		return nil, fmt.Errorf(c.GetRequiredMessage(input.ColumnName, input.RowNum))
	}

	var convertedData interface{}
	var err error
	switch input.ExpectedType {
	case "string":
		convertedData = input.CellData
	case "number":
		convertedData, err = strconv.Atoi(input.CellData) // Assuming the number is integer; use ParseFloat for float64
		if err != nil {
			return nil, fmt.Errorf(c.GetInvalidMessage(input.ColumnName, input.RowNum))
		}
	default:
		convertedData = input.CellData
	}

	if input.CustomMessageError != "" {
		return convertedData, fmt.Errorf(input.CustomMessageError)
	}

	return convertedData, nil
}

func (c *CheckValidImport) GetRequiredMessage(columnName string, rowNum int) string {
	return fmt.Sprintf("Dòng %d: %s là data bắt buộc", rowNum, columnName)
}

func (c *CheckValidImport) GetInvalidMessage(columnName string, rowNum int) string {
	return fmt.Sprintf("Dòng %d: %s không hợp lệ", rowNum, columnName)
}
