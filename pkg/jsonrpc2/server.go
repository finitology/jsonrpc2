package jsonrpc2

import (
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

	req, err := ParseRequest(body)
	if err != nil {
		resp := NewErrorResponse(nil, ErrInvalidRequest.WithData(err.Error()))
		s.writeResponse(w, resp)
		return
	}

	handler := s.router.Get(req.Method)
	if handler == nil {
		resp := NewErrorResponse(req.ID, ErrMethodNotFound)
		s.writeResponse(w, resp)
		return
	}

	result, rpcErr := handler(req)

	// Notification → no response
	if req.IsNotification() {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var resp *Response
	if rpcErr != nil {
		resp = NewErrorResponse(req.ID, rpcErr)
	} else {
		resp = NewSuccess(req.ID, result)
	}

	s.writeResponse(w, resp)
}

func (s *Server) writeResponse(w http.ResponseWriter, resp *Response) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := resp.Marshal()
	w.Write(data)
}
