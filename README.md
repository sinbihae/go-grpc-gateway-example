```
generate-user-proto:
	protoc -I=. \
	    --go_out . \
	    --go_opt paths=source_relative \
	    --go-grpc_out . --go-grpc_opt paths=source_relative \
	    protos/v1/user/user.proto
```

```
generate-post-proto:
	protoc -I=. \
	    --go_out . \
	    --go_opt paths=source_relative \
	    --go-grpc_out . --go-grpc_opt paths=source_relative \
	    protos/v1/post/post.proto
```
```
generate-user-v2-proto:
	protoc -I=. \
	    --go_out . \
        --go_opt paths=source_relative \
        --go-grpc_out . \
        --go-grpc_opt paths=source_relative \
	    protos/v2/user/user.proto
```
```
generate-user-v2-gateway-proto:
	protoc -I . \
		-I./grpc-gateway/third_party/googleapis/ \
	    --grpc-gateway_out . \
	    --grpc-gateway_opt logtostderr=true \
	    --grpc-gateway_opt paths=source_relative \
	    protos/v2/user/user.proto
```
```
go run simple-grpc-gateway/server/main.go
```
```
go run simple-grpc-gateway/grpc-gateway/main.go
```
# go-grpc-gateway-example
