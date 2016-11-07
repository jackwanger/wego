all: test

test:
	@go test

update_dict:
	@curl -LO https://github.com/downloads/wear/harmonious_dictionary/dictionaries.zip
	@unzip -o dictionaries.zip -d assets/
	@find assets/ -type f -print0 | xargs -0 sed -i 's/$$/ 2/'
	@go generate -x
	@rm -rf dictionaries.zip assets/

generate_dict:
	@go generate -x
