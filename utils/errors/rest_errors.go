package errors

import "net/http"

type RestErr struct {
	Message []string `json:"message"`
	Code    int      `json:"code"`
	Error   string   `json:"error"`
}

func NewBadRequest(message []string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Error:   "Bad request",
	}
}
func NewNotFound(message []string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message []string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

//NewAuthError - error handling for password rules, showed in slices
func PassError(message []string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusInternalServerError,
		Error:   "password_error",
	}
}
