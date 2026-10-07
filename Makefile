.PHONY: run build test vet fmt lint

run: 
	go run ./cmd/api
build: 
	go build -o bin/api ./cmd/api
test: 
	go test ./...
vet: 
	go vet ./...
fmt: 
	go fmt ./...

lint:
	golangci-lint run ./...