package api

import (
	"github.com/buddhimaaushan/mini_bank/db/sqlc"
	"github.com/buddhimaaushan/mini_bank/utils"
	"github.com/go-playground/validator/v10"
)

var validAccountStatus validator.Func = func(fl validator.FieldLevel) bool {
	if _, ok := fl.Field().Interface().(sqlc.Status); ok {
		return true
	}
	return false

}

var validEmailTypes validator.Func = func(fl validator.FieldLevel) bool {
	if email, ok := fl.Field().Interface().(string); ok {
		return utils.IsAValidEmailType(email)

	}
	return false
}
