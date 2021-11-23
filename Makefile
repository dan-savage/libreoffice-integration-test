install: build
	@sh -c "'$(CURDIR)/scripts/install.sh'"
test: 
	@sh -c "'$(CURDIR)/scripts/test.sh'"
integration:
	@sh -c "'$(CURDIR)/scripts/integration_test.sh'"