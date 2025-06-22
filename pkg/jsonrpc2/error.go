// Copyright 2025 Ganesh Anantwar (finitology)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
