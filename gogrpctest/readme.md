# gRPC Protocol Buffer 

## Install the dep package
$ brew install dep

## Install the Protobuf package

$ brew install protobuf

## To Generate protobuf file from protofile
### be in the folder that contains the xxx.proto file and run the below command
protoc --go_out=. testservice.proto