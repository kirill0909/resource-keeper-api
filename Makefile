.PHONY:
.SILENT:

build:
	go build -o .bin/main cmd/main.go

run: build
	./run.sh

