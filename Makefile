ifeq ($(VERSION),)
	VERSION := $(shell git describe --tags || git rev-parse --short HEAD)
endif

all: test

generate:
	@go generate -x ./...

release: generate
	@go build -ldflags "-s -w \
-X main.env=release \
-X main.buildstamp=`date -u '+%Y-%m-%dT%H:%M:%SZ'` \
-X main.githash=`git rev-parse HEAD` \
-X main.version=${VERSION}" -o wego

test:
	@go test

update_dict:
	@curl -LO https://github.com/downloads/wear/harmonious_dictionary/dictionaries.zip
	@unzip -o dictionaries.zip -d dict/assets/
	@rm -f dictionaries.zip
	@find dict/assets/ -type f -print0 | xargs -0 sed -i 's/$$/ 2/'
	@go generate -x ./...
