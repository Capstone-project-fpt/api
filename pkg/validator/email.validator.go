package validator

import (
	"regexp"
)

func IsValidFptEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@fpt\.edu\.vn$`)
	return re.MatchString(email)
}