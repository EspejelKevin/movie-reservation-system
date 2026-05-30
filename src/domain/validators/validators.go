package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	var (
		hasMinLen  = len(password) >= 8
		hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLower   = regexp.MustCompile(`[a-z]`).MatchString(password)
		hasNumber  = regexp.MustCompile(`[0-9]`).MatchString(password)
		hasSpecial = regexp.MustCompile(`[!@#~$%^&*()_+{}|:<>?=\-\[\]\\;',./]`).MatchString(password)
	)

	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func ValidateRole(regexRole string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		role := fl.Field().String()
		return regexp.MustCompile(regexRole).MatchString(role)
	}
}
