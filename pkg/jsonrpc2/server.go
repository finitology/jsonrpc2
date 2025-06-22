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
	"io"
	"net/http"
)

// Server wraps a router and handles HTTP JSON-RPC 2.0 requests.
type Server struct {
	router *Router
}

// NewServer creates a new Server.
func NewServer(router *Router) *Server {
	return &Server{router: router}
}

// ServeHTTP implements http.Handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost || r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read body", http.StatusInternalServerError)
		return
	}

	var jsonRaw json.RawMessage
	if err := json.Unmarshal(body, &jsonRaw); err != nil {
		s.writeResponse(w, NewErrorResponse(nil, ErrParse.WithData(err.Error())))
		return
	}

	// Batch request detection
	if len(jsonRaw) > 0 && jsonRaw[0] == '[' {
		var batch []json.RawMessage
		if err := json.Unmarshal(jsonRaw, &batch); err != nil {
			s.writeResponse(w, NewErrorResponse(nil, ErrInvalidRequest.WithData("invalid batch format")))
			return
		}

		if len(batch) == 0 {
			s.writeResponse(w, NewErrorResponse(nil, ErrInvalidRequest.WithData("empty batch")))
			return
		}

		var responses []json.RawMessage
		for _, item := range batch {
			resp := s.handleSingle(item)
			if resp != nil {
				data, err := resp.Marshal()
				if err == nil {
					responses = append(responses, data)
				}
			}
		}

		if len(responses) > 0 {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("[" + joinRawMessages(responses) + "]"))
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
		return // ✅ Prevent single-request handling
	}

	// Single request
	resp := s.handleSingle(jsonRaw)
	if resp != nil {
		s.writeResponse(w, resp)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// handleSingle parses and handles a single JSON-RPC request (not batch).
func (s *Server) handleSingle(data []byte) *Response {
	req, err := ParseRequest(data)
	if err != nil {
		return NewErrorResponse(nil, ErrInvalidRequest.WithData(err.Error()))
	}

	handler := s.router.Get(req.Method)
	if handler == nil {
		return NewErrorResponse(req.ID, ErrMethodNotFound)
	}

	result, rpcErr := handler(req)

	if req.IsNotification() {
		return nil
	}

	if rpcErr != nil {
		return NewErrorResponse(req.ID, rpcErr)
	}
	return NewSuccess(req.ID, result)
}

// writeResponse marshals and writes a response to the HTTP client.
func (s *Server) writeResponse(w http.ResponseWriter, resp *Response) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := resp.Marshal()
	w.Write(data)
}

// joinRawMessages joins multiple JSON RawMessages into a single JSON array string.
func joinRawMessages(msgs []json.RawMessage) string {
	var out []byte
	for i, m := range msgs {
		if i > 0 {
			out = append(out, ',')
		}
		out = append(out, m...)
	}
	return string(out)
}
