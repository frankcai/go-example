.PHONY: proto
MODULE="github.com/frankcai/go-example/microservices/proto"

proto:
	protoc --proto_path=. --go_out=. --go_opt=module=${MODULE} --go-grpc_out=. --go-grpc_opt=module=${MODULE} --micro_out=. ./*.proto
	mv ./${MODULE}/document/* ./document/
	rm -r ./github.com