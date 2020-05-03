# Makefile
APPNAME =`basename ${PWD}`
PACKAGE_DIRS=`go list -e ./... | egrep -v "binary_output_dir|.git|mocks"`

.PHONY: build run test protos
build:
	@go build
run: build
	@./$(APPNAME)
test:
	@go vet $(PACKAGE_DIRS)
	@go test $(PACKAGE_DIRS) -race -coverprofile=cover.out -covermode=atomic
dep:
	@glide up -v
protos:
	@protoc --proto_path=protobuf/schema \
			-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--go_out=plugins=grpc:protobuf \
			--grpc-gateway_out=logtostderr=true:protobuf \
			protobuf/schema/stock/*.proto