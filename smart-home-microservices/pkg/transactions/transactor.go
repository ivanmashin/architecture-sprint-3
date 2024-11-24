package transactions

import "context"

type Transactor[T any] interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context, repo T) error) error
}
