.PHONY: format
format:
	@go fmt ./...

.PHONY: build
build:
	@go build -o glox glox.go