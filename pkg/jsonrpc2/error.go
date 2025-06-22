package jsonrpc2

import (
	"encoding/json"
	"fmt"
)

// Error represents a JSON-RPC 2.0 error object.
type Error struct {
	Code    int64           `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data,omitempty"`
}

// Error implements the error interface
func (e *Error) Error() string {
	return fmt.Sprintf("jsonrpc2 error: code=%d, message=%q", e.Code, e.Message)
}

// WithData attaches custom error data
func (e *Error) WithData(data any) *Error {
	raw, err := json.Marshal(data)
	if err != nil {
		return e
	}
	e.Data = raw
	return e
}

// Predefined JSON-RPC 2.0 standard errors
var (
	ErrParse          = &Error{Code: -32700, Message: "Parse error"}
	ErrInvalidRequest = &Error{Code: -32600, Message: "Invalid Request"}
	ErrMethodNotFound = &Error{Code: -32601, Message: "Method not found"}
	ErrInvalidParams  = &Error{Code: -32602, Message: "Invalid params"}
	ErrInternal       = &Error{Code: -32603, Message: "Internal error"}
)

// NewError creates a custom JSON-RPC 2.0 error
func NewError(code int64, message string) *Error {
	return &Error{Code: code, Message: message}
}
