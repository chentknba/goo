package grpc

import (
    "goo/endpoint"

    "google.golang.org/grpc"

    "goo/examples/hss"
    grpctransport "goo/transport/grpc"
)

func New(conn *grpc.ClientConn) hss.Service {
    var serveEndpoint endpoint.Endpoint
    {
        serveEndpoint = grpctransport.NewClient(conn).Endpoint()
    }
}
