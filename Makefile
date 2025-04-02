.PHONY: test fmt tidy

test:
	go test ./... -v

fmt:
	go fmt ./...

tidy:
	go mod tidy
