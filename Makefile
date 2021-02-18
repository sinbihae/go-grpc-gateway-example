# Makefile

generate-proto:
	protoc -I=. \
	    --go_out . \
        --go_opt paths=source_relative \
        --go-grpc_out . \
        --go-grpc_opt paths=source_relative \
	    protos/v2/user/user.proto

generate-gateway-proto:
	protoc -I . \
		-I./grpc-gateway/third_party/googleapis/ \
	    --grpc-gateway_out . \
	    --grpc-gateway_opt logtostderr=true \
	    --grpc-gateway_opt paths=source_relative \
	    protos/v2/user/user.proto
