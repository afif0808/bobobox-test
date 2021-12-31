package customerrors

import "net/http"

type customError struct {
	Type    customErrorType
	Message string
	Err     error
}

type customErrorType string

func (cerr customError) Error() string {
	return cerr.Message + " : " + cerr.Err.Error()
}

const ErrTypeNotFound customErrorType = "not found"

func NewCustomError(msg string, err error, typ customErrorType) error {
	return customError{
		Type:    typ,
		Message: msg,
		Err:     err,
	}
}

func IsType(err error, typ customErrorType) bool {
	cerr, ok := err.(customError)
	if !ok || cerr.Type != typ {
		return false
	}
	return true
}

func ErrorHTTPCode(err error) int {
	cerr, ok := err.(customError)
	if !ok {
		return http.StatusInternalServerError
	}
	switch cerr.Type {
	case ErrTypeNotFound:
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
