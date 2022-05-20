/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/5/5-16:37
File: makefile.template.go
*/

package template


var MakefileTemplate = `export GO111MODULE:=on

BINARY_NAME:=crud
GO_FILES:=$(shell find . -name "*.go" -type f)

.PHONY: clean
clean:
	go clean && rm -f $(BINARY_NAME)

.PHONY: run
run:
	swag init && go run main.go

.PHONY: build
build:
	go build -o $(BINARY_NAME)

.PHONY: fmt
fmt:
	swag fmt && swag init
	go mod tidy
	gofmt -s -w $(GO_FILES)

.PHONY: vet
vet:
	go vet

.PHONY: install
install:
	go get -u github.com/swaggo/swag/cmd/swag
	# 1.16 or newer can use go install
	go get github.com/swaggo/swag
	go get github.com/gomodule/redigo

`
