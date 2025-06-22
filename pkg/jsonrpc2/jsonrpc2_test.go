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
	"testing"
)

func TestRouter_RegisterAndGet(t *testing.T) {
	r := NewRouter()
	err := r.Register("ping", func(req *Request) (any, *Error) {
		return "pong", nil
	})
	if err != nil {
		t.Fatalf("failed to register method: %v", err)
	}

	handler := r.Get("ping")
	if handler == nil {
		t.Fatal("handler not found after registration")
	}

	resp, rpcErr := handler(&Request{Method: "ping", JSONRPC: Version, ID: &ID{Str: strPtr("1")}})
	if rpcErr != nil {
		t.Fatalf("unexpected error: %v", rpcErr)
	}
	if resp != "pong" {
		t.Fatalf("expected pong, got %v", resp)
	}
}

func strPtr(s string) *string {
	return &s
}

func TestParseInvalidRequest(t *testing.T) {
	_, err := ParseRequest([]byte(`{}`))
	if err == nil {
		t.Error("expected error for invalid request")
	}

	_, err = ParseRequest([]byte(`invalid-json`))
	if err == nil {
		t.Error("expected error for bad JSON")
	}
}
