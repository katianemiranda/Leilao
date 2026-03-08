package internal_error

import "errors"

type InternalError struct {
	Message string
	Err     error
}

func (e *InternalError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     errors.New("not_found"),
	}
}

func NewInternalServerError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     errors.New("internal_server_error"),
	}
}

func NewBadRequestError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     errors.New("bad_request"),
	}
}
