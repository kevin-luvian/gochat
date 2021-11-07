build:
	go build -o bin/server main.go

run: build
	./bin/server

watch:
	ulimit -n 1000
	reflex -s -r '\.go$$' make run