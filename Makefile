## simple makefile to log workflow
.PHONY: all test clean build install

GOFLAGS ?= $(GOFLAGS:-v)

all: install

install: clean
	(cd pkg/my && $(CC) -shared -o libmy.so lib.c)
	@go get $(GOFLAGS) ./...

clean:
	@go clean $(GOFLAGS) -i ./...
	/bin/rm -f pkg/my/*.so
