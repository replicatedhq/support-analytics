SHELL := /bin/bash
.PHONY: clean deps install run test build swagger-spec shell all graphs force touch
.PRECIOUS: yaml/%.graph

SRC = $(shell find . -name "*.go")
YAMLS = $(wildcard yaml/*.yaml)
GRAPHS = $(YAMLS:.yaml=.png)

all: build test

clean:
	rm -f ./bin/support

deps:
	govendor install +local

run:
	./bin/support

test:
	go test -v `go list ./... | grep -v /vendor/`

build: bin/support

linux-build:
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -o ./bin/support .

install:
	cp bin/support /usr/local/bin

t:
	go test -v github.com/replicatedcom/support-analytics/bundle/replicated

graphs: $(GRAPHS)

force: touch build

touch:
	touch main.go

bin/support: $(SRC)
	mkdir -p bin
	go build -o ./bin/support .

yaml/%.graph: yaml/%.yaml bin/support
	bin/support events --yaml $< > $@

yaml/%.png: yaml/%.graph
	dot -Tpng $< > $@
