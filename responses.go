package rest

import (
	"encoding/json"
	"fmt"
)

type Resp struct {
	Resp interface{} `json:"response"`
	Err  error       `json:"error,omitempty"`
}

func (l *Resp) Error() string {
	b, err := json.Marshal(l)
	if err != nil {
		return "ErrResp: JSON marshaling error"
	}
	return fmt.Sprintf("%s", b)
}

type Meta struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

type MetaResp struct {
	Meta Meta        `json:"meta"`
	Body interface{} `json:"body"`
}

func (l *MetaResp) Error() string {
	b, err := json.Marshal(l)
	if err != nil {
		return "ErrResp: JSON marshaling error"
	}
	return fmt.Sprintf("%s", b)
}

type ListResp struct {
	Items []interface{} `json:"items"`
	Total int           `json:"total"`
	Err   error         `json:"error,omitempty"`
}

func (l *ListResp) AddItem(item interface{}) {
	l.Items = append(l.Items, item)
}

func (l *ListResp) Error() string {
	b, err := json.Marshal(l)
	if err != nil {
		return "ErrResp: JSON marshaling error"
	}
	return fmt.Sprintf("%s", b)
}

type ErrMeta struct {
	ErrCode    int    `json:"error_code"`
	ErrMessage string `json:"error_message"`
}

type ErrResp struct {
	Meta ErrMeta `json:"meta"`
}

func (e ErrResp) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return "ErrResp: JSON marshaling error"
	}
	return fmt.Sprintf("%s", b)
}

type ErrListResp struct {
	Meta   ErrMeta `json:"meta"`
	Errors []error `json:"errors"`
}

func (e ErrListResp) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return "ErrFieldResp: JSON marshaling error"
	}
	return fmt.Sprintf("%s", b)
}

func (e *ErrListResp) HasErrors() bool {
	return len(e.Errors) > 0
}

func (e *ErrListResp) AddError(err error) {
	e.Errors = append(e.Errors, err)
}

type ErrFieldResp struct {
	Meta   ErrMeta    `json:"meta"`
	Fields []ErrField `json:"fields,omitempty"`
	Errors []string   `json:"errors,omitempty"`
}

func (e ErrFieldResp) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return "ErrFieldResp: JSON marshaling error"
	}
	return fmt.Sprintf("%s", b)
}

func (e *ErrFieldResp) HasErrors() bool {
	return len(e.Fields) > 0
}

func (e *ErrFieldResp) Add(msg string) {
	e.Errors = append(e.Errors, msg)
}

func (e *ErrFieldResp) AddError(field string, code int, msg string) {
	for _, element := range e.Fields {
		if element.Field == field {
			element.AddError(code, msg)
			return
		}
	}

	f := ErrField{
		Field: field,
	}
	f.AddError(code, msg)
	e.AddField(f)
}

func (e *ErrFieldResp) AddField(field ErrField) {
	e.Fields = append(e.Fields, field)
}

type ErrField struct {
	Field string           `json:"field"`
	Errs  []ErrFieldObject `json:"errors"`
}

func (e *ErrField) AddError(code int, message string) {
	e.Errs = append(e.Errs, ErrFieldObject{
		Code:    code,
		Message: message,
	})
}

type ErrFieldObject struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_message"`
}
