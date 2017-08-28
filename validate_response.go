package rest

import (
	"github.com/go-ozzo/ozzo-validation"
	"strings"
)

func ValidateResponse(errors error) error {
	switch errors.(type) {
	case validation.Errors:
		errResponse := ErrFieldResp{
			Meta: ErrFieldRespMeta{
				ErrCode: 400,
			},
		}
		for k, v := range errors.(validation.Errors) {
			errResponse.AddError(strings.ToLower(k), 400, v.Error())
		}
		if errResponse.HasErrors() {
			return errResponse
		}
		return nil
	default:
		return errors
	}
}
