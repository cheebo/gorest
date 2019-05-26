package rest

import (
	"reflect"
)

func ValidateResponse(errors error, code int) error {
	errResponse := ErrFieldResp{
		Meta: ErrMeta{
			ErrCode: code,
		},
		Fields: []ErrField{},
	}

	v := reflect.ValueOf(errors)
	switch v.Kind() {
	// map[string]error
	case reflect.Map:
		if v.Type().Key().Kind() != reflect.String {
			break
		}
		if !v.Type().Elem().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			break
		}

		for _, k := range v.MapKeys() {
			errResponse.AddError(k.String(), 0, v.MapIndex(k).Interface().(error).Error())
		}
	// []error
	case reflect.Slice:
		if v.Len() == 0 {
			break
		}
		hasName := false
		if v.Index(0).Kind() == reflect.Interface {
			if v.Index(0).Elem().Kind() == reflect.Struct {
				if v.Index(0).Elem().FieldByName("Name").IsValid() {
					hasName = true
				}
			}
		}
		for i := 0; i < v.Len(); i++ {
			e := v.Index(i)
			if hasName {
				errResponse.AddError(
					e.Elem().FieldByName("Name").String(),
					0,
					e.Interface().(error).Error(),
				)
			} else {
				errResponse.Add(e.Interface().(error).Error())
			}
		}
	}

	return errResponse
}
