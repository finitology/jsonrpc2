package main

import (
	"log"
	"net/http"

	"github.com/finitology/jsonrpc2/pkg/jsonrpc2"
)

func main() {
	mux := jsonrpc2.NewRouter()
	mux.Register("ping", func(req *jsonrpc2.Request) (any, *jsonrpc2.Error) {
		return "pong", nil
	})

	srv := jsonrpc2.NewServer(mux)
	log.Println("JSON-RPC server listening on :8080")
	http.ListenAndServe(":8080", srv)
}
