package validators

import (
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

// ValidateCoolTitle returns true when the field value contains the word "cool".
func ValidateCoolTitle(fl validator.FieldLevel) bool {
	return strings.Contains(fl.Field().String(), "Cool")
}
