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
