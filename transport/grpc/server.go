package grpc

import (
    "context"

    "goo/endpoint"
)

type Handler interface {
    ServeGRPC(ctx context.Context, request interface{}) error
}

type Server struct {
    e endpoint.Endpoint
}

func NewServer(
    e endpoint.Endpoint) *Server{
    s := &Server{
        e : e,
    }

    return s
}

func (s *Server)ServeGRPC(ctx context.Context, request interface{}) error {
    return s.e(ctx, request)
}
