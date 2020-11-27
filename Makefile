BUILD_FILES = $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}}\
{{end}}' ./...)

RALLY_CLI_VERSION ?= $(shell git describe --tags 2>/dev/null || git rev-parse --short HEAD)
DATE_FMT = +%Y-%m-%d

BUILD_DATE ?= $(shell date "$(DATE_FMT)")

GO_LDFLAGS := -X github.com/psyllos/rally-cli/internal/build.Version=$(RALLY_CLI_VERSION) $(GO_LDFLAGS)
GO_LDFLAGS := -X github.com/psyllos/rally-cli/internal/build.Date=$(BUILD_DATE) $(GO_LDFLAGS)

bin/rally: $(BUILD_FILES)
	@go build -trimpath -ldflags "$(GO_LDFLAGS)" -o "$@" ./cmd/rally/

clean:
	rm -rf ./bin
.PHONY: clean

test:
	go test ./...
.PHONY: test
