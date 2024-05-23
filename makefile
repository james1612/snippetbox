# Variables
GO := go
BINARY_NAME := myapp
SOURCE_DIR := ./cmd/web

#################

run:
	$(GO) run $(SOURCE_DIR)

fmt:
	$(GO) fmt ./...