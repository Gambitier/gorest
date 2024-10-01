# Variables
SWAG_CMD=swag

# ANSI color codes
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[1;33m
BLUE=\033[0;34m
MAGENTA=\033[0;35m
CYAN=\033[0;36m
WHITE=\033[1;37m
NC=\033[0m # No Color

# Targets

.PHONY: all
all: help

.PHONY: help
help:
	@echo "Usage:"
	@echo "  make run                 Run the application with hot reload"
	@echo "  make build               Build source"
	@echo "  make install-tools       Install all required tools"
	@echo "  make swagger             Generate Swagger documentation"
	@echo "  make clean               Clean generated files"
	@echo "  make build-log           Generate Makefile build log"

.PHONY: build
build: 
	go mod tidy
	go build

.PHONY: install-tools
install-tools:
	go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: run
run:
	go run main.go

.PHONY: swagger
swagger:
	go generate ./...

.PHONY: clean
clean:
	@echo "Cleaning generated files..."
	rm -rf _apidocs
	rm -rf gorest

.PHONY: build-log
build-log:
	make --dry-run --always-make --keep-going --print-directory > Makefile-build.log