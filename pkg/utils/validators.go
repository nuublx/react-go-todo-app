package utils

import (
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/nuublx/react-go-todo-app/types"
)

func isValidUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	return len(username) >= 3 && len(username) <= 33
}

func isValidEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	reg := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return reg.MatchString(email)
}

func isValidPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	Low := regexp.MustCompile(`[a-z]`)
	Up := regexp.MustCompile(`[A-Z]`)
	Num := regexp.MustCompile(`[0-9]`)
	Spec := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)

	if !Low.MatchString(password) || !Up.MatchString(password) ||
		!Num.MatchString(password) || !Spec.MatchString(password) ||
		len(password) < 8 {
		return false
	}
	return true
}

func RegisterRequestValidator(model *types.RegisterRequest) error {
	validate := validator.New()
	validate.RegisterValidation("username", isValidUsername)
	validate.RegisterValidation("email", isValidEmail)
	validate.RegisterValidation("password", isValidPassword)
	validate.RegisterValidation("confirmpassword", isValidPassword)
	if err := validate.Struct(model); err != nil {
		return err
	}
	if model.Password != model.ConfirmPassword {
		var err = errors.New("passwords do not match")
		return err
	}

	return nil
}
