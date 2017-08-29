APP_NAME = xchange
GOBIN = $(GOPATH)/bin

go-build:
	GOOS=linux GOARCH=amd64 go build -o ./build/linux_amd64_$(APP_NAME) .
