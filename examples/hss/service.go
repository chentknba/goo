package hss

import (
    "context"
    "log"
    //"errors"
)

type Service interface {
    Serve(ctx context.Context, request []byte) error
}

type baseService struct {}

func (s baseService) Serve(ctx context.Context, request []byte) error {
    // handle the request
    log.Printf("recv: %s", string(request))
    return nil
}

func NewBasicService() Service {
    return baseService{}
}
