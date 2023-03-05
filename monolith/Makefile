.PHONY: proto
MODULE="github.com/frankcai/go-example/monolith"
PROTO_PATH="./proto"

proto:
	protoc --proto_path=${PROTO_PATH} --go_out=${PROTO_PATH} --go_opt=module=${MODULE} --go-grpc_out=${PROTO_PATH} --go-grpc_opt=module=${MODULE} --micro_out=${PROTO_PATH} ${PROTO_PATH}/*.proto
	mv ${PROTO_PATH}/${MODULE}/* ${PROTO_PATH}/
	rm -r ${PROTO_PATH}/github.com

build:
	go get -u ./...
	go build -o ./bin/monolith
	docker build --tag monolith .