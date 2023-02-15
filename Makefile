PLATFORMS := darwin-amd64 darwin-arm64 linux-amd64 linux-arm
TARGETS := $(PLATFORMS:%=build/%/mdtablefmt)
.PHONY: build test clean lint

build: lint test $(TARGETS)

clean:
	@echo '==> Cleaning'
	rm -rf -- build

lint: *.go
	@echo '==> Linting'
	go fmt
	go vet
	staticcheck

test: *.go
	@echo '==> Testing'
	go test

build/%: *.go
	@echo '==> Building $@'
	OUTPUT_FILE="$@"; \
	PLATFORM="$${OUTPUT_FILE##build/}"; PLATFORM="$${PLATFORM%%/*}"; \
	export GOOS="$${PLATFORM%%-*}" GOARCH="$${PLATFORM##*-}"; \
	go build -o "$${OUTPUT_FILE}"
