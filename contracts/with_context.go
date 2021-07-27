package contracts

import "context"

type WithContext interface {
	SetContext(ctx context.Context) interface{}
}
