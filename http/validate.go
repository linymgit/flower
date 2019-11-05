package http

import (
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

func ValidateReq(req interface{}) error {
	err := validate.Struct(req)
	var errMsg string
	if err != nil {
		vErrs := err.(validator.ValidationErrors)
		for _, fErr := range vErrs {
			errMsg = fmt.Sprintf("invalid request, field: %s, rule: %s, param: %s", fErr.Field(), fErr.Tag(), fErr.Param())
			break
		}
		if errMsg == "" {
			errMsg = "invalid request"
		}
		return errors.New(errMsg)
	}
	return nil
}
