package hss

import (
    "context"
    "goo/endpoint"
)

type Endpoints struct {
    ServeEndpoint endpoint.Endpoint
}

func (eps Endpoints) Serve(ctx context.Context, request []byte) error {
    return eps.ServeEndpoint(ctx, request)
}

func MakeServeEndpoint(s Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface {}) error {
        sreq := request.([]byte)
        return s.Serve(ctx, sreq)
    }
}
