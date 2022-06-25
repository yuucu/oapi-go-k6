

install:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

generate:
	go generate ./...
