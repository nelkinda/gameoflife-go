.PHONY: all
all: fmt test lint

.PHONY: fmt
fmt:
	go fmt

.PHONY: test
test: coverage.out

coverage.out: $(wildcard *.go)
	go test -test.coverprofile=coverage.out

.PHONY:
html: coverage.out
	go tool cover -html=coverage.out

.PHONY: lint
lint:
	golint

.PHONY: checkUpdates
checkUpdates:
	go list -u -m -f '{{if and .Update (not .Indirect)}}{{.}}{{end}}' all

.PHONY: clean
clean::
	$(RM) coverage.out
clean::
	go clean
