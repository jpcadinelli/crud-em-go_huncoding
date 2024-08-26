package rest_err

import "net/http"

type RestErr struct {
	Mensage string   `json:"mensage"`
	Err     string   `json:"err"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Mensage
}

func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Mensage: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(mensage string) *RestErr {
	return &RestErr{
		Mensage: mensage,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(mensage string, causes []Causes) *RestErr {
	return &RestErr{
		Mensage: mensage,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(mensage string) *RestErr {
	return &RestErr{
		Mensage: mensage,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(mensage string) *RestErr {
	return &RestErr{
		Mensage: mensage,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(mensage string) *RestErr {
	return &RestErr{
		Mensage: mensage,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
