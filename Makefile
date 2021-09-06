
#GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
#	go get github.com/micro/micro/v3/cmd/protoc-gen-micro
#	go get github.com/micro/micro/v3/cmd/protoc-gen-openapi

.PHONY: api
.PHONY: proto
proto:
	#protoc --proto_path=. /api --go_out=:. api/nextclipper.proto
	protoc --proto_path=.  --go_out=:. api/nextclipper.proto

.PHONY: build
build:
	go build -o helloworld *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t helloworld:latest
