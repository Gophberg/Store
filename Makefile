.PHONY: all build clean
.SILENT:

filename = "Store"

build:
	go build -o ./build/$(filename)__linux-x86_64 ./cmd/Store/main.go

	#GOOS=linux GOARCH=amd64 CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=0 go build -ldflags "-s -H linux" -o ./build/$(filename)__linux-x86_64 ./cmd/Store
	#GOOS=linux GOARCH=386 CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=0 go build -ldflags "-s -H linux" -o ./build/$(filename)__linux-x86 ./cmd/Store

run: build
	./build/$(filename)__linux-x86_64

test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := run
