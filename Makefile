.PHONY: test

GO := $(shell which go)

test: tests
tests:
	$(GO) test -v -failfast -cover ./lib/...

lint: linter
linter:
	goimports -w lib

install: setup
setup:
	$(GO) get -u golang.org/x/tools/cmd/goimports
