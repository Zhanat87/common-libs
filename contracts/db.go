package contracts

import "context"

type DB interface {
	BeginTransaction(ctx context.Context) (interface{}, error)
}
