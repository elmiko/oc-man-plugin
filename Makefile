.PHONY: clean tidy vendor

build:
	go build -o bin/oc-man

clean:
	rm bin/*

tidy:
	go mod tidy

vendor: tidy
	go mod vendor
