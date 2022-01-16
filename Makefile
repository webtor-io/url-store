protoc:
	protoc -I proto/ proto/url-store.proto --go_out=plugins=grpc:proto

build:
	go build .