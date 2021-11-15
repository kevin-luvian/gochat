install:
	go get ./...

build:
	go build -o bin/server main.go

test:
	GOFLAGS="-count=1" go test ./...

run: build
	./bin/server

watch:
	ulimit -n 1000
	reflex -s -r '\.go$$' make run