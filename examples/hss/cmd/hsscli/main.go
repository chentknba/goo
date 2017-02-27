package main

import (
    "context"
    //"io"
    "log"

    "time"

    pb "goo/examples/hss/pb"
    "google.golang.org/grpc"
)

const (
    port = ":9001"
)

func runClient() {
    conn, err := grpc.Dial(port, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to dial: %v", err)
    }

    defer conn.Close()

    client := pb.NewHssClient(conn)

    stream, err := client.Serve(context.Background())
    if err != nil {
        log.Fatalf("%v.Serve(_) = _, %v",client, err)
    }

    go func() {
        tick := time.Tick(500 * time.Millisecond)
        for _ = range(tick) {
            str := []byte("123")

            snd := &pb.Request{str}

            if err := stream.Send(snd); err != nil {
                log.Fatalf("failed to send msg: %v", err)
            }
        }
    }()
}

func main(){
    runClient()

    select{}
}
