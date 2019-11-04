GOFMT_FILES?=$$(find . -name '*.go')

bin/scraper: clean .vendor test
	 go build -o bin/scraper ./scraper

test: .fmtcheck
	go test -v ./...

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

.PHONY: test clean run

default: bin/scraper

.DEFAULT_GOAL := default
