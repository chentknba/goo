package grpc

import (
	//"context"

	//"goo/endpoint"

	"google.golang.org/grpc"
)

type Client struct {
	client *grpc.ClientConn
}

//func NewClient(cc *grpc.ClientConn) *Client {
//	return &Client{cc}
//}

func NewClient(addr string) (*Client, error) {
    conn, err := grpc.Dial(addr, grpc.WithInsecure())
    if err != nil {
        return &Client{conn}, err
    }

    return &Client{conn}
}
