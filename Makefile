APP_NAME = example
PKG = ./pkg/jsonrpc2

all: build

build:
	go build -o bin/$(APP_NAME) cmd/example/main.go

run:
	go run cmd/example/main.go

test:
	go test $(PKG)/... -v -cover

fmt:
	go fmt $(PKG)/...

lint:
	go vet $(PKG)/...

clean:
	rm -rf bin/

.PHONY: all build run test fmt lint clean
