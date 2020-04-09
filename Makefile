
INPUT=cmd/starter/main.go
OUTPUT=bin/starter
PKGS := $(shell go list ./... | grep -v /vendor)

.PHONY: $(OUTPUT)
$(OUTPUT):
	go build -o $(OUTPUT) $(INPUT)

.PHONY: build
build: $(OUTPUT)

.PHONY: run
run: build
	$(OUTPUT)

.PHONY: test
test:
	go test $(PKGS)
