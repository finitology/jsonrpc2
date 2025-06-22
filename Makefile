test:
	go test ./...

fmt:
	go fmt ./...

run:
	go run ./cmd/example/main.go

.PHONY: test fmt run
