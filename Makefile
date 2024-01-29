BIN ?= $(CURDIR)/bin
TMP ?= $(CURDIR)/.tmp


.PHONY: test
test:
	go test ./...


.PHONY: generate-schema
generate-schema:
	GOBIN=$(BIN) go install github.com/mprot/mprotc@latest
	mkdir -p $(TMP)
	curl -L -o $(TMP)/schema.mprot https://raw.githubusercontent.com/liblxn/lxn/main/schema.mprot
	$(BIN)/mprotc go --out internal/lxn/ --root $(TMP)/ schema.mprot
	rm -r $(TMP)
