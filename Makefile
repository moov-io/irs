PLATFORM=$(shell uname -s | tr '[:upper:]' '[:lower:]')
VERSION := $(shell grep -Eo '(v[0-9]+[\.][0-9]+[\.][0-9]+(-[a-zA-Z0-9]*)?)' version.go)

USERID := $(shell id -u $$USER)
GROUPID:= $(shell id -g $$USER)

build: irs

irs:
	pkger
	go build -o ${PWD}/bin/irs cmd/irs/*

run: irs
	./bin/irs

test: services build
	go test -cover ./...
	rm -rf cmd/irs/output

services:
	-docker-compose up -d --force-recreate

install:
	go install github.com/markbates/pkger/cmd/pkger
	git checkout LICENSE

.PHONY: check
check: build services
ifeq ($(OS),Windows_NT)
	@echo "Skipping checks on Windows, currently unsupported."
else
	@wget -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	./lint-project.sh
	@rm -rf cmd/irs/output
endif

dist: clean build
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=1 GOOS=windows go build -o bin/irs.exe github.com/moov-io/irs/cmd/irs
else
	CGO_ENABLED=1 GOOS=$(PLATFORM) go build -o bin/irs-$(PLATFORM)-amd64 github.com/moov-io/irs/cmd/irs
endif

docker: install
	pkger
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o ${PWD}/bin/.docker/irs cmd/irs/*
	docker build --pull -t moov/irs:$(VERSION) -f Dockerfile .
	docker tag moov/irs:$(VERSION) moov/irs:latest

docker-run:
	docker run -v ${PWD}/configs:/configs --env APP_CONFIG="/configs/config.yml" -it --rm moov/irs:$(VERSION)

docker-push:
	docker push moov/irs:$(VERSION)
	docker push moov/irs:latest

.PHONY: clean
clean:
ifeq ($(OS),Windows_NT)
	@echo "Skipping cleanup on Windows, currently unsupported."
else
	@rm -rf cover.out coverage.txt misspell* staticcheck*
	@rm -rf ./bin/ openapi-generator-cli-*.jar irs.db ./storage/ lint-project.sh
endif

.PHONY: cover-test cover-web
cover-test:
	go test -coverprofile=cover.out ./...
cover-web:
	go tool cover -html=cover.out

# Generate the go code from the public and internal api's
openapitools:
	docker run --rm \
		-u $(USERID):$(GROUPID) \
		-e OPENAPI_GENERATOR_VERSION='4.2.0' \
		-v ${PWD}:/local openapitools/openapi-generator-cli batch -- /local/.openapi-generator/client-generator-config.yml

# From https://github.com/genuinetools/img
.PHONY: AUTHORS
AUTHORS:
	@$(file >$@,# This file lists all individuals having contributed content to the repository.)
	@$(file >>$@,# For how it is generated, see `make AUTHORS`.)
	@echo "$(shell git log --format='\n%aN <%aE>' | LC_ALL=C.UTF-8 sort -uf)" >> $@
