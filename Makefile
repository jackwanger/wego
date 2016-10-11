all: test

generate:
	@go generate -x ./...

release: generate
	@go build -ldflags "-X main.env=release -X main.buildstamp=`date -u '+%Y-%m-%dT%H:%M:%SZ'` -X main.githash=`git rev-parse HEAD`" -o wego

test:
	@go test

update_dict:
	@curl -LO https://github.com/downloads/wear/harmonious_dictionary/dictionaries.zip
	@unzip -o dictionaries.zip -d dict/assets/
	@rm -f dictionaries.zip
	@find dict/assets/ -type f -print0 | xargs -0 sed -i 's/$$/ 2/'
	@go generate -x ./...

.PHONY: generate release update_dict
