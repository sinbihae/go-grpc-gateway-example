package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	networkpb "github.com/sinbihae/go-grpc-gateway-example/protos/network"
)

const (
	portNumber           = "9000"
	gRPCServerPortNumber = "9001"
)

// WithIncomingHeaderMatcher gateway에 들어온 요청에 대해 특정 헤더만 허용하는 옵션이다.
// Client로부터 요청이 왔을 때 여기에 정의된 헤더를 포함하는 경우에만 gRPC server에 요청을 전달한다.
// 소무낮로 작성해야하군....
func CustomMatcher(key string) (string, bool) {
	switch key {
	case "x-koscom-access-key":
		return key, true
	case "x-koscom-secret-key":
		return key, true
	default:
		return key, true
		// return runtime.DefaultHeaderMatcher(key)
	}
}

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(CustomMatcher),
	)
	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	if err := networkpb.RegisterVpcHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:"+gRPCServerPortNumber,
		options,
	); err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	log.Printf("start HTTP server on %s port", portNumber)
	if err := http.ListenAndServe(":"+portNumber, mux); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
