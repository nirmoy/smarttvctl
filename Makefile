SHELL := /usr/bin/env bash
CWD := $(shell pwd)
BIN := smarttvctl 

SOURCES := $(shell find  . -name '*.go')

.PHONY: clean

all: $(BIN)

$(BIN): $(SOURCES)
	GO111MODULE=on go build 

clean:
	rm -f $(BIN)
