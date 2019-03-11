SHELL := /usr/bin/env bash
CWD := $(shell pwd)
BIN := smarttvctl 

SOURCES := $(shell find  . -name '*.go')

.PHONY: clean

all: $(BIN)

$(BIN): $(SOURCES)
	go build -o $(BIN) main.go

clean:
	rm -f $(BIN)
