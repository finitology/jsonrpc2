package jsonrpc2

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBatchRequest_Mixed(t *testing.T) {
	router := NewRouter()
	_ = router.Register("add", func(req *Request) (any, *Error) {
		var params struct {
			A int `json:"a"`
			B int `json:"b"`
		}
		if err := json.Unmarshal(req.Params, &params); err != nil {
			return nil, ErrInvalidParams
		}
		return params.A + params.B, nil
	})

	server := NewServer(router)

	batch := []any{
		map[string]any{
			"jsonrpc": "2.0",
			"method":  "add",
			"params":  map[string]any{"a": 1, "b": 2},
			"id":      1,
		},
		map[string]any{
			"jsonrpc": "2.0",
			"method":  "nonexistent",
			"id":      2,
		},
		map[string]any{
			"jsonrpc": "2.0",
			"method":  "add",
			"params":  map[string]any{"a": 5, "b": 7},
			"id":      3,
		},
		map[string]any{ // notification, no response
			"jsonrpc": "2.0",
			"method":  "add",
			"params":  map[string]any{"a": 99, "b": 1},
		},
	}

	data, _ := json.Marshal(batch)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", rec.Code)
	}

	var resp []Response
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}

	if len(resp) != 3 {
		t.Fatalf("expected 3 responses (1 notification skipped), got %d", len(resp))
	}

	for _, r := range resp {
		if r.Error != nil && r.ID != nil && r.ID.Num != nil && *r.ID.Num == 2 {
			if r.Error.Code != -32601 {
				t.Errorf("expected method not found error for id 2")
			}
		} else if r.Result != nil && r.ID != nil && r.ID.Num != nil {
			// Expecting results for id 1 and 3
			if *r.ID.Num == 1 && r.Result != float64(3) {
				t.Errorf("expected 3 for id 1, got %v", r.Result)
			}
			if *r.ID.Num == 3 && r.Result != float64(12) {
				t.Errorf("expected 12 for id 3, got %v", r.Result)
			}
		}
	}
}
