install:
	go get ./...

build:
	go build -o bin/server main.go

test:
	GOFLAGS="-count=1" go test ./...

run: build
	./bin/server

run_swag: swagger build run

lint:
	golint ./...

swagger:
	swag init

watch:
	ulimit -n 1000
	reflex -s -R '^docs/' -r '\.go$$' make run
