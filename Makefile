generate:
	#protoc proto/your_service.proto --go_out=plugins=grpc:proto
	#protoc --grpc-gateway_out=logtostderr=true proto/your_service.proto
	protoc \
    		-I. \
    		-I=${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway\@v1.14.6/third_party/googleapis \
    		--go_out=plugins=grpc:proto \
    		--grpc-gateway_out=logtostderr=true:./proto \
    		--swagger_out=logtostderr=true:. \
    		proto/*.proto;