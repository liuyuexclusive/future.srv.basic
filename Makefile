
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/user/user.proto
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/role/role.proto
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/message/message.proto

.PHONY: build
build: proto

	go build -o basic-srv main.go plugin.go
	

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t basic-srv:latest
