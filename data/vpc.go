package data

import (
	networkpb "github.com/sinbihae/go-grpc-gateway-example/protos/network"
)

var Vpcs = []*networkpb.VpcMessage{
	{
		Id:          "4783",
		Name:        "dev",
		CidrBlock:   "192.168.0.0/16",
		Status:      "RUN",
		Region:      "FKR",
		CreatedDate: "1607994910000",
	},
	{
		Id:          "4578",
		Name:        "vpc-hjy",
		CidrBlock:   "172.30.0.0/16",
		Status:      "RUN",
		Region:      "FKR",
		CreatedDate: "1606178416000",
	},
	{
		Id:          "2688",
		Name:        "jihoon",
		CidrBlock:   "172.16.0.0/16",
		Status:      "RUN",
		Region:      "FKR",
		CreatedDate: "1586159375000",
	},
}
