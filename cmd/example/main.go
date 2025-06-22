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

package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	jsonrpc2 "github.com/finitology/jsonrpc2/pkg/jsonrpc2"
)

func add(req *jsonrpc2.Request) (any, *jsonrpc2.Error) {
	var p struct {
		A int `json:"a"`
		B int `json:"b"`
	}
	if err := json.Unmarshal(req.Params, &p); err != nil {
		return nil, jsonrpc2.ErrInvalidParams
	}
	return p.A + p.B, nil
}

func ping(req *jsonrpc2.Request) (any, *jsonrpc2.Error) {
	log.Println("ping received")
	return "pong", nil
}

func main() {
	router := jsonrpc2.NewRouter()
	router.Register("ping", ping)

	router.Register("add", add)

	server := &http.Server{
		Addr:    ":8080",
		Handler: jsonrpc2.NewServer(router),
	}

	// Run server in goroutine
	go func() {
		log.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	// Wait for SIGINT or SIGTERM
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}

	log.Println("Server exited gracefully")
}
