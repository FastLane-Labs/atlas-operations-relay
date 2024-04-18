package relayerror

import (
	"encoding/json"
	"fmt"
)

var (
	// Errors used across different packages
	ErrServerInternal = NewError(1000, "server internal error")
)

type Error struct {
	Code    uint16 `json:"code"`
	Message string `json:"message"`
}

func NewError(code uint16, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func (e *Error) AddError(err error) *Error {
	return NewError(
		e.Code,
		fmt.Sprintf("%s: %s", e.Message, err.Error()),
	)
}

func (e *Error) AddMessage(message string) *Error {
	return NewError(
		e.Code,
		fmt.Sprintf("%s: %s", e.Message, message),
	)
}

func (e *Error) Marshal() []byte {
	b, _ := json.Marshal(e)
	return b
}
