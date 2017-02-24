package main

import (
	"context"
	"encoding/json"
	//"fmt"
	"io"
	"log"
	//"net"

    "time"

	pb "goo/examples/helloworld/helloworld"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
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

func runClient(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewHelloWorldClient(conn)

    stream, err := client.Hello(context.Background())
    if err != nil {
        log.Fatalf("%v.Hello(_) = _, %v", client, err)
    }

    waitc := make(chan struct{})

    // recv
    go func() {
        for {
            in, err := stream.Recv()
            if err == io.EOF {
                close(waitc)
                return
            }

            if err != nil {
                log.Fatalf("failed to recv msg: %v", err)
            }

            var resp jsonResponse

            if err = json.Unmarshal(in.Payload, &resp); err != nil {
                log.Fatalf("failed to unmarshal msg: %v", err)
            }

            log.Printf("recv response, code:%d", resp.Code)
        }
    }()

    // send
    go func() {
        tick := time.Tick(10*time.Millisecond)
        for _ = range(tick) {
            req := jsonRequest{"func1", "damn"}

            r, err := json.Marshal(req)
            if err != nil {
                log.Fatalf("failed to marshal msg: %v", err)
            }

            snd := &pb.Request{r}
            if err := stream.Send(snd); err != nil {
                log.Fatalf("failed to send msg: %v", err)
            }
        }
    }()

    <-waitc
}

func main() {
    runClient(port1)
}
