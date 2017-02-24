package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"time"

	pb "goo/examples/helloworld/helloworld"
	"google.golang.org/grpc"
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

func runClient(port string) {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial:%v", port)
	}
	defer conn.Close()

	client := pb.NewHelloWorldClient(conn)

	stream, err := client.Hello(context.Background())
	if err != nil {
		log.Fatalf("%v.Hello() = _, %v", client, err)
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
				log.Fatalf("failed to receive msg : %v\n", err)
			}

			var resp jsonResponse
			if err := json.Unmarshal(in.Payload, &resp); err != nil {
				log.Fatalf("failed to unmarshal msg: %v\n", err)
			}

			log.Printf("recv from server, code:%d", resp.Code)
		}
	}()

	// send
	go func() {
		tick := time.Tick(50 * time.Millisecond)

		for _ = range tick {
			req := jsonRequest{"aaa", "bbb"}
			r, err := json.Marshal(req)
			if err != nil {
				log.Fatalf("marshal req err:%v\n", err)
			}

			snd := &pb.Request{r}
			if err := stream.Send(snd); err != nil {
				log.Fatalf("send msg err:%v\n", err)
			}

		}

		close(waitc)
	}()

	//select {}
	<-waitc
}

func main() {
	go runClient(port1)
	go runClient(port2)

	select {}
}
