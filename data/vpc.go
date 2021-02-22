package data

import (
	networkpb "github.com/sinbihae/go-grpc-gateway-example/protos/network"
)

var Vpcs = []*networkpb.VpcMessage{
	{
		VpcNo:         "4783",
		VpcName:       "dev",
		Ipv4CidrBlock: "192.168.0.0/16",
		VpcStatus:     "RUN",
		RegionCode:    "FKR",
		CreateDate:    "1607994910000",
	},
	{
		VpcNo:         "4578",
		VpcName:       "vpc-hjy",
		Ipv4CidrBlock: "172.30.0.0/16",
		VpcStatus:     "RUN",
		RegionCode:    "FKR",
		CreateDate:    "1606178416000",
	},
	{
		VpcNo:         "2688",
		VpcName:       "jihoon",
		Ipv4CidrBlock: "172.16.0.0/16",
		VpcStatus:     "RUN",
		RegionCode:    "FKR",
		CreateDate:    "1586159375000",
	},
}
