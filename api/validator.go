package api

import (
	"github.com/buddhimaaushan/mini_bank/utils"
	"github.com/go-playground/validator/v10"
)

var validAccountStatus validator.Func = func(fl validator.FieldLevel) bool {
	if status, ok := fl.Field().Interface().(string); ok {
		return utils.IsVerifiedAccStatus(status)
	}
	return false
}
