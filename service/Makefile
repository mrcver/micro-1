GOPATH:=$(shell go env GOPATH)

.PHONY: build docker
all: build docker

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w' -i -o user.service service.go

.PHONY: docker
docker:
	docker build . -t user-service:latest
