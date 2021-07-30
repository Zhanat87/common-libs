package contracts

import "context"

type Transaction interface {
	BeginTransaction(ctx context.Context) (interface{}, error)
}
