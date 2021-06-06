.PHONY: build
build:
	mkdir -p ./out
	go build -ldflags="-w -s" -o ./out/pbcli ./cmd/pbcli

.PHONY: test
test:
	stdout=$$(gofmt -l . 2>&1); \
	if [ "$$stdout" ]; then \
		exit 1; \
	fi
	go vet ./...
	gocyclo -over 10 $(shell find . -iname '*.go' -type f)
	staticcheck ./...
	go test -v -cover ./...

.PHONY: setup
setup:
	go get -u github.com/fzipp/gocyclo/cmd/gocyclo
	go get -u honnef.co/go/tools/cmd/staticcheck
