.PHONY: clean tidy vendor

all: gather build

build:
	go build -o bin/oc-man

clean:
	rm -r bin content

gather:
	python3 hack/gather.py

tidy:
	go mod tidy

vendor: tidy
	go mod vendor
