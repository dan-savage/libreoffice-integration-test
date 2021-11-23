install: build
	@sh -c "'$(CURDIR)/scripts/install.sh'"
test: 
	@sh -c "'$(CURDIR)/scripts/test.sh'"