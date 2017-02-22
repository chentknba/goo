package endpoint

import (
	"context"
)

type func(ctx context.Context, request interface{}) (interface{}, error)
