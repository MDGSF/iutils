package rspmsg

import (
	"fmt"
	"net/http"
)

const (
	CodeSuccess = "0"
	CodeError   = "1"

	CodeSuccessMsg = "success"
)

type ErrMsg struct {
	HTTPCode int         `json:"httpCode"` // http.StatusOK
	Code     string      `json:"code"`     // 0 - success, 1 - error
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
}

func (c *ErrMsg) Error() string {
	return fmt.Sprintf("HTTPCode: %d,code: %s, msg: %s",
		c.HTTPCode, c.Code, c.Message)
}

func NewErrMsg(HTTPCode int, code string, msg string) error {
	return &ErrMsg{HTTPCode: HTTPCode, Code: code, Message: msg}
}

func NewErr400(msg string) error {
	return &ErrMsg{
		HTTPCode: http.StatusBadRequest,
		Code:     CodeError,
		Message:  msg,
	}
}

func NewErr500(msg string) error {
	return &ErrMsg{
		HTTPCode: http.StatusInternalServerError,
		Code:     CodeError,
		Message:  http.StatusText(http.StatusInternalServerError),
	}
}

type RspMsg[T any] struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func RspMsgFromErrMsg(v any) RspMsg[any] {
	var resp RspMsg[any]
	switch data := v.(type) {
	case *ErrMsg:
		resp.Code = data.Code
		resp.Message = data.Message
	case ErrMsg:
		resp.Code = data.Code
		resp.Message = data.Message
	default:
		resp.Code = CodeSuccess
		resp.Message = CodeSuccessMsg
		resp.Data = v
	}
	return resp
}
