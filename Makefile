GOFMT_FILES?=$$(find . -name '*.go')
EQUALS := =

bin/scraper: .vendor test
	 go build -o bin/scraper ./src

test: .fmtcheck
	go test ./... -v -race -coverprofile$(EQUALS)coverage.txt -covermode$(EQUALS)atomic

clean:
	rm -rf ./bin

fmt: .tools
	goimports -w $(GOFMT_FILES)

.tools:
	go get golang.org/x/tools/cmd/goimports

.fmtcheck:
	@goimports -l $(GOFMT_FILES) | read; if [ $$? == 0 ]; then echo "gofmt check failed for:"; goimports -l $(GOFMT_FILES); exit 1; fi

.vendor:
	go mod download
	go mod tidy
	@touch .vendor

.PHONY: test clean run fmt

default: bin/scraper

.DEFAULT_GOAL := default
