.PHONY: build
build:
	go build -ldflags="-w -s" -o pbcli ./cmd

.PHONY: test
test:
	stdout=$$(gofmt -l . 2>&1); \
	if [ "$$stdout" ]; then \
		exit 1; \
	fi
	gocyclo -over 10 $(shell find . -iname '*.go' -type f)
	go test -v -cover ./...
	stdout=$$(golint ./... 2>&1); \
	if [ "$$stdout" ]; then \
		exit 1; \
	fi

.PHONY: setup
setup:
	go get -u github.com/fzipp/gocyclo/cmd/gocyclo
	go get -u golang.org/x/lint/golint
