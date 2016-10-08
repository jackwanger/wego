all: test

generate:
	@go generate -x

release: generate
	@go build

test:
	@go test

update_dict:
	@curl -LO https://github.com/downloads/wear/harmonious_dictionary/dictionaries.zip
	@unzip -o dictionaries.zip -d dict/
	@rm -f dictionaries.zip
	@find dict -type f -print0 | xargs -0 sed -i 's/$$/ 2/'
	@go generate -x ./...

.PHONY: generate release update_dict
