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
## tailwind dev
css-watch:
	@npx tailwindcss -i ./internal/static/views/css/input.css -o ./internal/static/views/css/output.css --watch
## dev
dev:
	@air
## tailwind build
