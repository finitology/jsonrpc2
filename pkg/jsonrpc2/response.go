package jsonrpc2

import (
	"encoding/json"
)

// Response represents a JSON-RPC 2.0 response object.
type Response struct {
	JSONRPC string `json:"jsonrpc"`
	Result  any    `json:"result,omitempty"`
	Error   *Error `json:"error,omitempty"`
	ID      *ID    `json:"id"`
}

// NewSuccess creates a successful response with result.
func NewSuccess(id *ID, result any) *Response {
	return &Response{
		JSONRPC: Version,
		Result:  result,
		ID:      id,
	}
}

// NewErrorResponse creates a failed response with error.
func NewErrorResponse(id *ID, err *Error) *Response {
	return &Response{
		JSONRPC: Version,
		Error:   err,
		ID:      id,
	}
}

// Marshal returns the response as JSON.
func (r *Response) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
