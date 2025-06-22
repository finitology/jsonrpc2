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
	"fmt"
	"sync"
)

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
func (r *Router) Register(method string, fn HandlerFunc) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.handlers[method]; exists {
		return fmt.Errorf("method already registered: %q", method)
	}

	r.handlers[method] = fn
	return nil
}

// Get returns the handler for a method, or nil.
func (r *Router) Get(method string) HandlerFunc {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.handlers[method]
}
