package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func NewValidator() *Validator {
	v := validator.New()
	v.RegisterValidation("bookIdValidation", validateBookId)

	return &Validator{
		validator: v,
	}
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func validateBookId(fl validator.FieldLevel) bool {
	bookId := fl.Field().String()
	idPattern := `^S\d{12}$`
	reg := regexp.MustCompile(idPattern)
	result := reg.MatchString(bookId)
	return result
}

func ValidateStockId(id string) bool {
	idPattern := `^S\d{12}$`
	reg := regexp.MustCompile(idPattern)
	return reg.MatchString(id)
}
