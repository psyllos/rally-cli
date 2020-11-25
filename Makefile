BUILD_FILES = $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}}\
{{end}}' ./...)

DATE_FMT = +%Y-%m-%d
BUILD_DATE ?= $(shell date "$(DATE_FMT)")

GO_LDFLAGS := -X github.com/psyllos/rally-cli/internal/build.Date=$(BUILD_DATE) $(GO_LDFLAGS)

bin/rally: $(BUILD_FILES)
	go build -trimpath -ldflags "$(GO_LDFLAGS)" -o "$@" ./cmd/rally
# .PHONY: bin/rally

clean:
	rm -rf ./bin
.PHONY: clean
