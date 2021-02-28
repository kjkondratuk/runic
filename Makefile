GO=go
MAKE=make

run:
	$(GO) run cmd/runic.go $(TEXT)

build:
	$(GO) build -o bin/runic cmd/runic.go

run-compiled: build
	./bin/runic $(TEXT)