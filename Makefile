all: deps test install-linux install-windows

install: install-linux install-windows

deps:
	go get -v -t -d

test:
	go test ./... -v -cover

install-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags "-X main.version=$(version)"

install-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -v -ldflags "-X main.version=$(version)"
