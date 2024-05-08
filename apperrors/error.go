package apperrors

import "errors"

type TodoAppError struct {
	ErrCode
	Message string
	Err     error
}

func (e *TodoAppError) Error() string {
	return e.Err.Error()
}

var ErrNoData = errors.New("no data found")
