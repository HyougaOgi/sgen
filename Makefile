.PHONY: all run clean format

.DEFAULT_GOAL := build

build:
	go build -o sgen ./cmd

run-test:
	./sgen test

run-clean:
	./sgen clean
clean:
	rm -f sgen

format:
	find . -type f -name '*.go' -print0 | xargs -0 -n 1 gofmt -w
