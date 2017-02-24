package main

import (
	//"context"
	"encoding/json"
	//"fmt"
	"io"
	"log"
	"net"

	pb "goo/examples/helloworld/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	//"google.golang.org/grpc/metadata"
)

const (
	port1 = ":9001"
	port2 = ":9002"
)

type jsonRequest struct {
	Method string `json:method`
	Param  string `json:param`
}

type jsonResponse struct {
	Code int `json:code`
}

type server struct{}

// Hello
func (s *server) Hello(stream pb.HelloWorld_HelloServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		var req jsonRequest
		if err := json.Unmarshal(in.Payload, &req); err != nil {
			return err
		}

		log.Printf("recv msg, method:%s, param:%s\n", req.Method, req.Param)

		resp := jsonResponse{123456700001}
		b, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("snd msg marshal err:%v", err)
		}

		snd := &pb.Response{b}
		if err := stream.Send(snd); err != nil {
			log.Fatalf("send msg err:%v", err)
		}
	}
}

func runServer(port string) {
	go func() {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("fail to listen: %v\n", err)
		}

		s := grpc.NewServer()
		pb.RegisterHelloWorldServer(s, &server{})

		reflection.Register(s)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v\n", err)
		}
	}()
}

func main() {
	runServer(port1)
	runServer(port2)

	select {}
}
