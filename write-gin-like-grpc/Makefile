.PHONY: all lint

lint:
	gofmt -s -w ./
	goimports -w ./
	golangci-lint run