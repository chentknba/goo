package hss

import (
    "context"

    "io"

	"log"

    "goo/examples/hss/pb"
    grpctransport "goo/transport/grpc"
)

func MakeGRPCServer(ctx context.Context, endpoints Endpoints) pb.HssServer {
    return &grpcServer{
        serve : grpctransport.NewServer(endpoints.ServeEndpoint),
    }
}

type grpcServer struct {
    serve grpctransport.Handler
}

func (s *grpcServer) Serve(stream pb.Hss_ServeServer) error {
    for {
        in, err := stream.Recv()
        if err == io.EOF {
            return nil
        }

        if err != nil {
            log.Printf("recv error, %v\n", err)
            return err
        }

        ctx := context.Background()
        r, err := DecodeGRPCServeRequest(ctx, in)
        s.serve.ServeGRPC(ctx, r)
    }
}

func DecodeGRPCServeRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
    req := grpcReq.(*pb.Request)
    return req.Payload, nil
}
