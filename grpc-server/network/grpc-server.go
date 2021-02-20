package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	networkpb "github.com/sinbihae/go-grpc-gateway-example/protos/network"
	// userpb "github.com/sinbihae/go-grpc-gateway-example/protos/v2/user"
)

const portNumber2 = "9001"

type vpcServer struct {
	networkpb.VpcServer
}

// GetUser returns user message by user_id
func (s *vpcServer) GetVpc(ctx context.Context, req *networkpb.GetVpcRequest) (*networkpb.GetVpcResponse, error) {
	// vpcNo := req.VpcNo

	var vpcMessage *networkpb.VpcMessage
	// for _, u := range data.UsersV2 {
	// 	if u.UserId != userID {
	// 		continue
	// 	}
	// 	vpcMessage = u
	// 	break
	// }

	return &networkpb.GetVpcResponse{
		VpcMessage: vpcMessage,
	}, nil
}

// ListUsers returns all user messages
func (s *vpcServer) ListVpcs(ctx context.Context, req *networkpb.ListVpcsRequest) (*networkpb.ListVpcsResponse, error) {
	vpcMessages := make([]*networkpb.VpcMessage, 3)
	// for i, u := range data.UsersV2 {
	// 	userMessages[i] = u
	// }

	return &networkpb.ListVpcsResponse{
		VpcMessages: vpcMessages,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber2)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	networkpb.RegisterVpcServer(grpcServer, &vpcServer{})

	log.Printf("start gRPC server on %s port", portNumber2)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
