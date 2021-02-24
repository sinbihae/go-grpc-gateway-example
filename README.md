


## grpc-gateway
gRPC server와 HTTP로 RESTful 하게 통신할 수 있게 해 준다. gRPC gateway도 protobuf를 기반으로 작동을 한다. Protobuf 컴파일러인 protoc 로 gRPC server에서 사용하기 위한 message들 및 서비스를 생성했던 것 처럼, protoc로 gRPC gateway에서 사용할 수 있는 message들과 서비스를 생성해주는 것이다. gRPC gateway에서는 protobuf 파일을 읽고 HTTP JSON을 protobuf로 변환한 뒤에 gRPC로 reverse proxy해준다.

![grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway/raw/master/docs/assets/images/architecture_introduction_diagram.svg)




### grpc-kotlin-client 구성

```
WEB ---->  Gateway ---->  API Server(Grpc-Kotlin-Client  ----> Grpc-Go-Server(NCP Client)) ----> NCP Server
```

- Grpc-Kotlin-Client: req/rep objec 생성.
- Grpc-Go-Server: grpc proto 정의.
- NCP client: Swagger 생성.



#### grpc-gateway 도입 구성
```
WEB ---->  Gateway ---->  API Server((Grpc-gateway) ----> Grpc-Go-Server(NCP Client)) ----> NCP Server
```

- Grpc-gateway: grpc proto 그대로 사용.
- Grpc-Go-Server: grpc proto 정의.
- NCP client: Swagger 생성.



#### proto 파일 작성 방법

```
  syntax = "proto3";

  package user;
+ 
+ import "google/api/annotations.proto";

  service User {
-   rpc GetUser(GetUserRequest) returns (GetUserResponse);
+   rpc GetUser(GetUserRequest) returns (GetUserResponse) {
+     option (google.api.http) = {
+       get: "/v1/users/{user_id}"
+     };
+   }
  }
```

#### grpc-gateway 확인 사항
1. method(GET, POST, PUT, DELETE, FETCH) 모두 정의 가능
2. Path Param, QueryString, Body 정의 가능
3. Custom Header 적용 가능
* 소스 코드 참고 

#### Grpc Server 구동

```
go run grpc-server/network/grpc-server.go
```


#### HTTP Server 구동

```
go run gateway-server/network/http-server.go
```
