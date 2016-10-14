ifeq ($(VERSION),)
	VERSION := $(shell git describe --tags || git rev-parse --short HEAD)
endif

all: test

release:
	@go build -ldflags "-s -w \
-X main.env=release \
-X main.buildstamp=`date -u '+%Y-%m-%dT%H:%M:%SZ'` \
-X main.githash=`git rev-parse HEAD` \
-X main.version=${VERSION}" -o wego

test:
	@go test
