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
