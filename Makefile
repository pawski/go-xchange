APP_NAME = xchange
GOBIN = $(GOPATH)/bin

go-build:
	go build -o ./cache/$(APP_NAME) .
