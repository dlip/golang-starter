
INPUT=cmd/starter/main.go
OUTPUT=bin/starter

.PHONY: $(OUTPUT)
$(OUTPUT):
	go build -o $(OUTPUT) $(INPUT)

.PHONY: build
build: $(OUTPUT)

.PHONY: run
run: build
	$(OUTPUT)
