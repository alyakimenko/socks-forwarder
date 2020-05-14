.PHONY: build

DIR_BIN = ./bin

default: build

build:
	go build -v -o ${DIR_BIN}/ferret \
	./cmd/...
