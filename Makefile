all: deps test build-linux build-windows

deps:
	go get -v -t -d

test:
	go test ./... -v -cover

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -v