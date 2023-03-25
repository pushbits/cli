OUT_DIR := ./out
GO_FILES := $(shell find . -type f \( -iname '*.go' \))

PBCLI_BUILD_VERSION ?= $(shell git describe --tags)
ifeq ($(PBCLI_BUILD_VERSION),)
	_ := $(error Cannot determine build version)
endif

BUILD_VERSION_FLAG := github.com/pushbits/cli/internal/buildconfig.Version=$(PBCLI_BUILD_VERSION)

.PHONY: build
build:
	mkdir -p $(OUT_DIR)
	go build -ldflags "-w -s -X $(BUILD_VERSION_FLAG)" -o $(OUT_DIR)/pbcli ./cmd/pbcli

.PHONY: clean
clean:
	rm -rf $(OUT_DIR)

.PHONY: test
test:
	stdout=$$(gofumpt -l . 2>&1); if [ "$$stdout" ]; then exit 1; fi
	go vet ./...
	gocyclo -over 10 $(GO_FILES)
	staticcheck ./...
	errcheck ./...
	go test -v -cover ./...
	gosec -exclude-dir=tests ./...
	govulncheck ./...
	@printf '\n%s\n' "> Test successful"

.PHONY: setup
setup:
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	go install github.com/kisielk/errcheck@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install mvdan.cc/gofumpt@latest

.PHONY: fmt
fmt:
	gofumpt -l -w .
