package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func notEmptyStringFunc(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[^\s-]`)
	return re.MatchString(fl.Field().String())
}

func nameThFunc(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[ก-๛]+`)
	return re.MatchString(fl.Field().String())
}

func nameEnFunc(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[A-Za-z]+`)
	return re.MatchString(fl.Field().String())
}
