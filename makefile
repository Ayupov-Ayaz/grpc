install-dependencies:
	go get -u google.golang.org/grpc
	go get -u google.golang.org/protobuf

generate:
	protoc -I api/proto --go_out=plugins=grpc:. api/proto/adder.proto

