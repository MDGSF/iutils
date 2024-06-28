package rspmsg

import (
	"net/http"
	"reflect"
	"testing"
)

func TestErrMsg_Error(t *testing.T) {
	errMsg := &ErrMsg{
		HTTPCode: 404,
		Code:     "NOT_FOUND",
		Message:  "Resource not found",
	}

	expected := "HTTPCode: 404,code: NOT_FOUND, msg: Resource not found"
	actual := errMsg.Error()

	if actual != expected {
		t.Errorf("Expected error message to be: %s, but got: %s", expected, actual)
	}
}

func TestNewErrMsg(t *testing.T) {
	tests := []struct {
		name     string
		HTTPCode int
		code     string
		msg      string
	}{
		{
			name:     "Error with HTTPCode 400",
			HTTPCode: 400,
			code:     "ERR_400",
			msg:      "Bad Request",
		},
		{
			name:     "Error with HTTPCode 404",
			HTTPCode: 404,
			code:     "ERR_404",
			msg:      "Not Found",
		},
		{
			name:     "Error with HTTPCode 500",
			HTTPCode: 500,
			code:     "ERR_500",
			msg:      "Internal Server Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewErrMsg(tt.HTTPCode, tt.code, tt.msg)
			if err == nil {
				t.Errorf("NewErrMsg() error = %v, wantErr %v", err, tt.msg)
				return
			}
			myErr, ok := err.(*ErrMsg)
			if !ok {
				t.Errorf("NewErrMsg() error = %v, not of type *ErrMsg", err)
				return
			}
			if myErr.HTTPCode != tt.HTTPCode {
				t.Errorf("NewErrMsg() HTTPCode = %v, want %v", myErr.HTTPCode, tt.HTTPCode)
			}
			if myErr.Code != tt.code {
				t.Errorf("NewErrMsg() Code = %v, want %v", myErr.Code, tt.code)
			}
			if myErr.Message != tt.msg {
				t.Errorf("NewErrMsg() Message = %v, want %v", myErr.Message, tt.msg)
			}
		})
	}
}

func TestNewErr400(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{
			name:    "Error with message",
			msg:     "Invalid request",
			wantErr: true,
		},
		{
			name:    "Error without message",
			msg:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewErr400(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewErr400() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				myErr, ok := err.(*ErrMsg)
				if !ok {
					t.Errorf("NewErr400() returned wrong type, got %T, want *ErrMsg", err)
				}
				if myErr.HTTPCode != http.StatusBadRequest {
					t.Errorf("NewErr400() HTTPCode = %v, want %v", myErr.HTTPCode, http.StatusBadRequest)
				}
				if myErr.Code != CodeError {
					t.Errorf("NewErr400() Code = %v, want %v", myErr.Code, CodeError)
				}
				if myErr.Message != tt.msg {
					t.Errorf("NewErr400() Message = %v, want %v", myErr.Message, tt.msg)
				}
			}
		})
	}
}

func TestNewErr401(t *testing.T) {
	msg := "Unauthorized"
	err := NewErr401(msg)

	if err == nil {
		t.Errorf("Expected error to be not nil")
	}

	errMsg, ok := err.(*ErrMsg)
	if !ok {
		t.Errorf("Expected error to be of type *ErrMsg")
	}

	if errMsg.HTTPCode != http.StatusUnauthorized {
		t.Errorf("Expected HTTPCode to be %d, but got %d", http.StatusUnauthorized, errMsg.HTTPCode)
	}

	if errMsg.Code != CodeError {
		t.Errorf("Expected Code to be %v, but got %v", CodeError, errMsg.Code)
	}

	if errMsg.Message != msg {
		t.Errorf("Expected Message to be %s, but got %s", msg, errMsg.Message)
	}
}

func TestNewErr403(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{
			name:    "Error message is empty",
			msg:     "",
			wantErr: true,
		},
		{
			name:    "Error message is not empty",
			msg:     "Access denied",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewErr403(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewErr403() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				myErr, ok := err.(*ErrMsg)
				if !ok {
					t.Errorf("NewErr403() returned wrong type, got %T, want *ErrMsg", err)
				}
				if myErr.HTTPCode != http.StatusForbidden {
					t.Errorf("NewErr403() HTTPCode = %v, want %v", myErr.HTTPCode, http.StatusForbidden)
				}
				if myErr.Code != CodeError {
					t.Errorf("NewErr403() Code = %v, want %v", myErr.Code, CodeError)
				}
				if myErr.Message != tt.msg {
					t.Errorf("NewErr403() Message = %v, want %v", myErr.Message, tt.msg)
				}
			}
		})
	}
}

func TestNewErr500(t *testing.T) {
	err := NewErr500()

	if err == nil {
		t.Errorf("NewErr500() should return a non-nil error")
	}

	if _, ok := err.(*ErrMsg); !ok {
		t.Errorf("NewErr500() should return a pointer to ErrMsg")
	}

	errMsg, ok := err.(*ErrMsg)
	if !ok {
		t.Fatalf("Failed to assert type of error")
	}

	if errMsg.HTTPCode != http.StatusInternalServerError {
		t.Errorf("Expected HTTPCode to be %d, got %d", http.StatusInternalServerError, errMsg.HTTPCode)
	}

	if errMsg.Code != CodeError {
		t.Errorf("Expected Code to be %v, got %v", CodeError, errMsg.Code)
	}

	if errMsg.Message != http.StatusText(http.StatusInternalServerError) {
		t.Errorf("Expected Message to be %s, got %s", http.StatusText(http.StatusInternalServerError), errMsg.Message)
	}
}

func TestNewRspMsg(t *testing.T) {
	tests := []struct {
		name         string
		lrsp         any
		err          any
		wantResp     RspMsg[any]
		wantHTTPCode int
	}{
		{
			name:         "Success",
			lrsp:         "data",
			err:          nil,
			wantResp:     RspMsg[any]{Code: CodeSuccess, Message: CodeSuccessMsg, Data: "data"},
			wantHTTPCode: http.StatusOK,
		},
		{
			name:         "Error with ErrMsg pointer",
			lrsp:         "data",
			err:          &ErrMsg{Code: "400", Message: "Bad Request", HTTPCode: http.StatusBadRequest},
			wantResp:     RspMsg[any]{Code: "400", Message: "Bad Request", Data: "data"},
			wantHTTPCode: http.StatusBadRequest,
		},
		{
			name:         "Error with ErrMsg value",
			lrsp:         "data",
			err:          ErrMsg{Code: "500", Message: "Internal Server Error", HTTPCode: http.StatusInternalServerError},
			wantResp:     RspMsg[any]{Code: "500", Message: "Internal Server Error", Data: "data"},
			wantHTTPCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, gotHTTPCode := NewRspMsg(tt.lrsp, tt.err)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NewRspMsg() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotHTTPCode != tt.wantHTTPCode {
				t.Errorf("NewRspMsg() gotHTTPCode = %v, want %v", gotHTTPCode, tt.wantHTTPCode)
			}
		})
	}
}
