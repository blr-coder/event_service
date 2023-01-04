package usecase_errors

import "fmt"

type ValidationErr struct {
	ErrMessage string
}

func NewValidationErr(err error) *ValidationErr {
	return &ValidationErr{ErrMessage: err.Error()}
}

func (e *ValidationErr) Error() string {
	return fmt.Sprintf("input validation error, %s", e.ErrMessage)
}
