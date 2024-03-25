gen-proto:
	protoc -I ./pkg/protos/proto ./pkg/protos/proto/sso/sso.proto --go_out=./pkg/protos/gen/go --go_opt=paths=source_relative --go-grpc_out=./pkg/protos/gen/go/ --go-grpc_opt=paths=source_relative

build:
	go build cmd/sso/main.go

run:
	go run cmd/sso/main.go --config=./config/local.yaml