TESTALL = $$(go list ./...)
TESTE2E = ./tests/...
GOTEST = GODEBUG=cgocheck=0 go test --count=1

test:
	$(GOTEST) --timeout 5m $(TESTALL)

test-integration:
	$(GOTEST) --timeout 5m $(TESTE2E)
