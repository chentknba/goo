package main

import (
	"context"
	"log"
	"net"

	"goo/endpoint"

	"goo/examples/hss"
	"goo/examples/hss/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
    addr = "localhost:9001"
)

func main() {
	var service hss.Service
	{
		service = hss.NewBasicService()
	}

	var serveEndpoint endpoint.Endpoint
	{
		serveEndpoint = hss.MakeServeEndpoint(service)
	}

	eps := hss.Endpoints{ServeEndpoint: serveEndpoint}

	go func() {
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("failed to listen, err: %v\n", err)
		}

		srv := hss.MakeGRPCServer(context.Background(), eps)
		s := grpc.NewServer()

		pb.RegisterHssServer(s, srv)

		reflection.Register(s)

		log.Printf("start gprc server.\n")

		if err := s.Serve(ln); err != nil {
			log.Fatalf("failed to serve:%v\n", err)
		}
	}()

	select {}
}
