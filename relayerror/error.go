package relayerror

import (
	"encoding/json"
	"fmt"
)

var (
	ErrComputeUserOpHash = NewError(3000, "can't compute user operation hash")
	ErrServerInternal    = NewError(3001, "server internal error")
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
	e.Message = fmt.Sprintf("%s: %s", e.Message, err.Error())
	return e
}

func (e *Error) Marshal() []byte {
	b, _ := json.Marshal(e)
	return b
}
