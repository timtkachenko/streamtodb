install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2

lint:
	golangci-lint run

test:
	go test -cover -race ./...

service:
	go run cmd/main.go cmd/config.go ./ports.json
