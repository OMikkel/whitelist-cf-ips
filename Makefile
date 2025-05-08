build:
	make build-x86-linux
	make build-arm-macos

build-x86-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/whitelist-cf-ips-x86-linux cmd/server/main.go

build-arm-macos:
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o bin/whitelist-cf-ips-arm-macos cmd/server/main.go

run:
	go run cmd/server/main.go
