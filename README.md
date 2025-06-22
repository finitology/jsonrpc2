# jsonrpc2

> A pure Go implementation of the JSON-RPC 2.0 specification. Minimal, fast, and dependency-free.


[![CI](https://github.com/finitology/jsonrpc2/actions/workflows/ci.yml/badge.svg)](https://github.com/finitology/jsonrpc2/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/finitology/jsonrpc2)](https://goreportcard.com/report/github.com/finitology/jsonrpc2)
[![Go Reference](https://pkg.go.dev/badge/github.com/finitology/jsonrpc2.svg)](https://pkg.go.dev/github.com/finitology/jsonrpc2)
[![Coverage Status](https://coveralls.io/repos/github/finitology/jsonrpc2/badge.svg?branch=main)](https://coveralls.io/github/finitology/jsonrpc2?branch=main)
[![Release](https://img.shields.io/github/v/release/finitology/jsonrpc2)](https://github.com/finitology/jsonrpc2/releases)


> Minimal, idiomatic, and spec-compliant [JSON-RPC 2.0](https://www.jsonrpc.org/specification) implementation in pure Go (no third-party deps).

## ✨ Features

- Fully compliant with JSON-RPC 2.0
- Request, response, and error formatting
- Supports batch requests and notifications
- Thread-safe method registration
- Works with standard `net/http`
- Lightweight and dependency-free

## 🚀 Installation

```bash
go get github.com/finitology/jsonrpc2
```

## 🔧 Usage
```
mux := jsonrpc2.NewRouter()
mux.Register("ping", func(req *jsonrpc2.Request) (any, *jsonrpc2.Error) {
    return "pong", nil
})

srv := jsonrpc2.NewServer(mux)
log.Fatal(http.ListenAndServe(":8080", srv))
```

## 📦 Example
```bash
curl -X POST http://localhost:8080 \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"ping","id":1}'

```

## 📂 Project Structure

```bash
jsonrpc2/
├── pkg/jsonrpc2     # Core library code
├── cmd/example      # Example app for testing
├── .github/workflows
├── Makefile

```

## 🛡 License

This project is licensed under the [Apache 2.0 License](./LICENSE).

---

