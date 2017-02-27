package endpoint

import (
	"context"
)

type Endpoint func(ctx context.Context, request interface{}) error
