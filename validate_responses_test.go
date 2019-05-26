package rest_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cheebo/gorest"

	"github.com/go-ozzo/ozzo-validation"

	"github.com/asaskevich/govalidator"
)

func TestValidateResponse_Slice(t *testing.T) {
	a := assert.New(t)

	errMsg := "fieldName is invalid"

	errs := ErrorSlice{
		errors.New(errMsg),
	}

	err := rest.ValidateResponse(errs, 400)

	exp := rest.ErrFieldResp{
		Meta: rest.ErrMeta{
			ErrCode:    400,
			ErrMessage: "",
		},
		Fields: []rest.ErrField{},
		Errors: []string{errMsg},
	}

	a.Equal(exp, err)
}

func TestValidateResponse_CustomErrStruct(t *testing.T) {
	a := assert.New(t)

	errField := "fieldName"
	errMsg := "fielName is invalid"

	errs := ErrorSlice{
		Err{
			Name: errField,
			Err:  errors.New(errMsg),
		},
	}

	err := rest.ValidateResponse(errs, 400)

	exp := rest.ErrFieldResp{
		Meta: rest.ErrMeta{
			ErrCode:    400,
			ErrMessage: "",
		},
		Fields: []rest.ErrField{
			rest.ErrField{
				Field: errField,
				Errs: []rest.ErrFieldObject{
					rest.ErrFieldObject{
						Message: errMsg,
					},
				},
			},
		},
	}

	a.Equal(exp, err)
}

func TestValidateResponse_Govalidator(t *testing.T) {
	a := assert.New(t)

	errField := "fieldName"
	errMsg := "fielName is invalid"

	errs := govalidator.Errors{
		govalidator.Error{
			Name:                     errField,
			CustomErrorMessageExists: true,
			Err:                      errors.New(errMsg),
		},
	}

	err := rest.ValidateResponse(errs, 400)

	exp := rest.ErrFieldResp{
		Meta: rest.ErrMeta{
			ErrCode:    400,
			ErrMessage: "",
		},
		Fields: []rest.ErrField{
			rest.ErrField{
				Field: errField,
				Errs: []rest.ErrFieldObject{
					rest.ErrFieldObject{
						Message: errMsg,
					},
				},
			},
		},
	}

	a.Equal(exp, err)
}

func TestValidateResponse_Ozzo(t *testing.T) {
	a := assert.New(t)

	errField := "fieldName"
	errMsg := "fielName is invalid"

	errs := validation.Errors{
		errField: errors.New(errMsg),
	}

	err := rest.ValidateResponse(errs, 400)

	exp := rest.ErrFieldResp{
		Meta: rest.ErrMeta{
			ErrCode:    400,
			ErrMessage: "",
		},
		Fields: []rest.ErrField{
			rest.ErrField{
				Field: errField,
				Errs: []rest.ErrFieldObject{
					rest.ErrFieldObject{
						Message: errMsg,
					},
				},
			},
		},
	}

	a.Equal(exp, err)
}

type ErrorSlice []error

func (es ErrorSlice) Error() string {
	var errs []string
	for _, e := range es {
		errs = append(errs, e.Error())
	}
	return strings.Join(errs, ";")
}

type Err struct {
	Name string
	Err  error
}

func (e Err) Error() string {
	return fmt.Sprintf("%s", e.Err.Error())
}
