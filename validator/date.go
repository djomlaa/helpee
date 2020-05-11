package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// ValidateDateOfBirth checks if date of birth is valid
func ValidateDateOfBirth(field validator.FieldLevel) bool {
	now := time.Now()
	var wayBack int64= -2208993840000 // 1900-01-01
	return wayBack < field.Field().Int() && field.Field().Int() < now.Unix()
}