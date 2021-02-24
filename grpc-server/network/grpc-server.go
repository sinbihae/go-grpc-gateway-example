package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sinbihae/go-grpc-gateway-example/data"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	networkpb "github.com/sinbihae/go-grpc-gateway-example/protos/network"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const portNumber2 = "9001"

type vpcServer struct {
	networkpb.VpcServer
}

func (s *vpcServer) GetVpc(ctx context.Context, req *networkpb.VpcSearch) (*networkpb.VpcsResponse, error) {
	Id := req.Id

	vpcMessages := make([]*networkpb.VpcMessage, 1)
	for _, u := range data.Vpcs {
		if u.Id == Id {
			vpcMessages[0] = u
			break
		}
	}

	return &networkpb.VpcsResponse{
		VpcMessages: vpcMessages,
	}, nil
}

func (s *vpcServer) ListVpcs(ctx context.Context, req *networkpb.VpcSearch) (*networkpb.VpcsResponse, error) {
	vpcMessages := make([]*networkpb.VpcMessage, 3)
	for i, u := range data.Vpcs {
		vpcMessages[i] = u
	}

	return &networkpb.VpcsResponse{
		VpcMessages: vpcMessages,
	}, nil
}

func (s *vpcServer) CreateVpc(ctx context.Context, req *networkpb.VpcMessage) (*networkpb.VpcsResponse, error) {

	fmt.Printf("req:%s", req)
	vpcMessages := make([]*networkpb.VpcMessage, 1)
	vpcMessages[0] = req

	return &networkpb.VpcsResponse{
		VpcMessages: vpcMessages,
	}, nil
}

func (s *vpcServer) UpdateVpc(ctx context.Context, req *networkpb.VpcMessage) (*networkpb.VpcsResponse, error) {

	fmt.Printf("req:%s", req)
	vpcMessages := make([]*networkpb.VpcMessage, 1)
	vpcMessages[0] = req

	return &networkpb.VpcsResponse{
		VpcMessages: vpcMessages,
	}, nil
}

func (s *vpcServer) DeleteVpc(ctx context.Context, req *networkpb.VpcSearch) (*emptypb.Empty, error) {
	Id := req.Id
	fmt.Printf("Id:%s", Id)

	koscomAccessKey := ""
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if key, ok := md["x-koscom-access-key"]; ok {
			koscomAccessKey = key[0]
		}
		fmt.Println(md)
	}

	fmt.Printf("key:%s", koscomAccessKey)

	return &emptypb.Empty{}, nil
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
