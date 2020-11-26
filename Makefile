APP_NAME = xchange
GOBIN = $(GOPATH)/bin

go-build:
	GOOS=linux GOARCH=amd64 go build -o ./build/linux_64_$(APP_NAME) ./cmd/xchange/
	GOOS=darwin GOARCH=amd64 go build -o ./build/darwin_64_$(APP_NAME) ./cmd/xchange/
