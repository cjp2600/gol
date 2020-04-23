build:
	protoc -I.  \
	-I$(GOPATH)/src   \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   \
	--plugin=protoc-gen-mongo=$(GOPATH)/bin/protoc-gen-bom \
	--mongo_out="generateCrud=true,gateway:." \
 	proto/gol.proto
 	
	protoc -I. \
	-I$(GOPATH)/src \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--plugin=protoc-gen-go=$(GOPATH)/bin/protoc-gen-go \
	--go_out=. \
	proto/gol.proto

	protoc -I. \
	-I$(GOPATH)/src \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--plugin=protoc-gen-go=$(GOPATH)/bin/protoc-gen-go \
	--proto_path=$(GOPATH)/src:. \
	--go_out=plugins=grpc:. \
	proto/gol.proto

	protoc -I. \
	-I$(GOPATH)/src \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--plugin=protoc-gen-grpc-gateway=$(GOPATH)/bin/protoc-gen-grpc-gateway \
	--grpc-gateway_out=logtostderr=true:. \
	proto/gol.proto

	protoc -I.  \
	-I$(GOPATH)/src   \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   \
	--plugin=protoc-gen-swagger=$(GOPATH)/bin/protoc-gen-swagger \
	--swagger_out=logtostderr=true:.  \
 	proto/gol.proto
