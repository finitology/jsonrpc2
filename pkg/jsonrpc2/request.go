package jsonrpc2

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Request represents a JSON-RPC 2.0 request object.
type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
	ID      *ID             `json:"id,omitempty"`
}

// IsNotification returns true if the request is a notification (no response required)
func (r *Request) IsNotification() bool {
	return r.ID == nil
}

// Validate checks whether the request adheres to JSON-RPC 2.0 spec
func (r *Request) Validate() error {
	if r.JSONRPC != Version {
		return fmt.Errorf("invalid jsonrpc version: got %q, want %q", r.JSONRPC, Version)
	}
	if r.Method == "" {
		return errors.New("method field is required")
	}
	return nil
}

// ParseRequest parses a JSON-encoded request into a Request struct
func ParseRequest(data []byte) (*Request, error) {
	var req Request
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, err
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &req, nil
}
