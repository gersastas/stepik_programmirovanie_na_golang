GOLANGCI_LINT := golangci-lint

.PHONY: lint

lint:
	$(GOLANGCI_LINT) run