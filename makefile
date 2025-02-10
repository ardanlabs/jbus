# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

run:
	go run api/services/sales/main.go | go run api/tooling/logfmt/main.go

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor
