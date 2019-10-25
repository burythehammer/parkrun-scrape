
bin/scraper: clean .vendor test
	go build -o bin/scraper -v

test:
	go test -v ./...

clean:
	go clean
	rm -f ./bin

.vendor:
	go mod download
	go mod tidy
	@touch .vendor

.PHONY: test clean run

default: bin/scraper

.DEFAULT_GOAL := default
