all: test

test:
	@go test

# require gnu sed
update_dict:
	@curl -LO https://github.com/downloads/wear/harmonious_dictionary/dictionaries.zip
	@unzip -o dictionaries.zip -d assets/
	@find assets/ -type f -print0 | xargs -0 sed 's/$/ 2/' -i
	@go generate -x
	@rm -rf dictionaries.zip assets/

generate_dict:
	@go generate -x
