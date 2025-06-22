package jsonrpc2

import "sync"

// HandlerFunc is a function that processes a JSON-RPC request.
type HandlerFunc func(*Request) (any, *Error)

// Router maps method names to handler functions.
type Router struct {
	mu       sync.RWMutex
	handlers map[string]HandlerFunc
}

// NewRouter creates a new Router instance.
func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
	}
}

// Register associates a method name with a handler.
func (r *Router) Register(method string, fn HandlerFunc) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.handlers[method] = fn
}

// Get returns the handler for a method, or nil.
func (r *Router) Get(method string) HandlerFunc {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.handlers[method]
}
