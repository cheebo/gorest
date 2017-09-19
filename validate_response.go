package rest

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-ozzo/ozzo-validation"
	"strings"
)

func ValidateResponse(errors error) error {
	switch errors.(type) {
	case govalidator.Errors:
		errResponse := ErrFieldResp{
			Meta: ErrFieldRespMeta{
				ErrCode: 400,
			},
		}
		for k, v := range errors.(govalidator.Errors) {
			e, ok := v.(govalidator.Error)
			if ok {
				errResponse.AddError(e.Name, 400, e.Error())
			} else {
				errResponse.AddError(string(k), 400, v.Error())
			}
		}
		if errResponse.HasErrors() {
			return errResponse
		}
		return nil
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
