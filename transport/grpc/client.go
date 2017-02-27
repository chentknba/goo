package grpc

import (
	//"context"

    //"log"

	//"goo/endpoint"

	"google.golang.org/grpc"
)

type Client struct {
	client *grpc.ClientConn
}

func NewClient(cc *grpc.ClientConn) *Client {
	return &Client{cc}
}
