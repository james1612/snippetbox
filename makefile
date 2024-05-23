# Variables
GO := go
BINARY_NAME := myapp
SOURCE_DIR := ./cmd/web

# Rules
run:
	$(GO) run $(SOURCE_DIR)
