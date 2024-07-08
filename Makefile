name := $(shell basename ${PWD})
PACKAGES := $(shell go list ./...)
.PHONY:
tidy:
	@echo " == executing go mod tidy.... =="
	@go mod tidy
## build
build :
	@echo "=== building the project ==="
	@go build -o  tmp/app
## vet 
vet :
	@echo "executing go vet"
	@go vet ./...
## dev
dev:
	@air
## tailwind build
