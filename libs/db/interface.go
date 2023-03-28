package db

//go:generate mockery --case underscore --name TxDB --with-expecter

import (
	"context"
)

type TxLevel string

const (
	Serializable    TxLevel = "serializable"
	RepeatableRead  TxLevel = "repeatable read"
	ReadCommitted   TxLevel = "read committed"
	ReadUncommitted TxLevel = "read uncommitted"
)

type RowsAffecter interface {
	RowsAffected() int64
}

type TxDB interface {
	InTx(ctx context.Context, lvl TxLevel, fx func(ctxTx context.Context) error) error
	Get(ctx context.Context, dst interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (RowsAffecter, error)
	Select(ctx context.Context, dst interface{}, query string, args ...interface{}) error
}
