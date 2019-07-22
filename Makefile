.PHONY: default clean checks test build

TAG_NAME := $(shell git tag -l --contains HEAD)
SHA := $(shell git rev-parse --short HEAD)
VERSION := $(if $(TAG_NAME),$(TAG_NAME),$(SHA))

BUILD_DATE := $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

export GO111MODULE=on

default: clean checks test build

test: clean
	go test -v -cover ./...

clean:
	rm -rf dist/ cover.out

dependencies:
	go mod vendor

build: clean
	@echo Version: $(VERSION) $(BUILD_DATE)
	go build -v -ldflags '-X "main.version=${VERSION}" -X "main.commit=${SHA}" -X "main.date=${BUILD_DATE}"' -o gh-check ./cmd

checks:
	golangci-lint run