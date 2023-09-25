VERSION ?= $(shell git describe --match 'v[0-9]*' --tags --always)

build:
	@go build -ldflags "-X main.version=$(VERSION)" ./cmd/app

version:
	@echo $(VERSION)

generate-docs:
	@swag init --parseDependency -g internal/app/api/v1/api.go

generate-di:
	@cd cmd/app && wire

run:
	@docker start go-api-database &> /dev/null || docker run -d \
	 --name go-api-database \
	 -p 27017:27017 mongo &> /dev/null
	go run -ldflags "-X main.version=$(VERSION)"  ./cmd/app/main.go ./cmd/app/wire_gen.go

docker:
	@docker build -t portfolio-api:$(VERSION)

test:
	@go test -v ./...

generate:
	@go generate ./...

tools:
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %