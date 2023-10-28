.PHONY:
.SILENT:
.DEFAULT_GOAL := test

test:
	go test --short -coverprofile=coverage.out -v ./api

	go tool cover -func=coverage.out | grep "total"