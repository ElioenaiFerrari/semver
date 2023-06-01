.PHONY: test

test:
	go test ./... -cover -coverprofile=coverage.out -v && go tool cover -html=coverage.out -o coverage.html