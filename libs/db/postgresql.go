package db

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/multierr"
)

type txkey string

const key = txkey("tx")

var _ TxDB = (*pgxPoolDB)(nil)

type pgxPoolDB struct {
	pool *pgxpool.Pool
}

type Querier interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
}

func NewPgxPoolDB(pool *pgxpool.Pool) *pgxPoolDB {
	return &pgxPoolDB{
		pool: pool,
	}
}

func (p *pgxPoolDB) InTx(ctx context.Context, lvl TxLevel, fx func(ctxTx context.Context) error) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "run transaction")
	defer span.Finish()

	tx, err := p.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.TxIsoLevel(lvl),
	})
	if err != nil {
		return err
	}

	err = fx(context.WithValue(ctx, key, tx))
	if err != nil {
		ext.Error.Set(span, true)
		span.SetTag("err_text", err)
		return multierr.Combine(err, tx.Rollback(ctx))
	}

	err = tx.Commit(ctx)
	if err != nil {
		return multierr.Combine(err, tx.Rollback(ctx))
	}

	return nil
}

func (p *pgxPoolDB) getQuerier(ctx context.Context) Querier {
	tx, ok := ctx.Value(key).(Querier)
	if ok && tx != nil {
		return tx
	}

	return p.pool
}

func (p *pgxPoolDB) Get(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "query get")
	defer span.Finish()
	span.SetTag("query", query)

	return pgxscan.Get(ctx, p.getQuerier(ctx), dst, query, args...)
}

func (p *pgxPoolDB) Exec(ctx context.Context, query string, args ...interface{}) (RowsAffecter, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "query exec")
	defer span.Finish()
	span.SetTag("query", query)

	rows, err := p.getQuerier(ctx).Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	rows.Close()
	return rows.CommandTag(), nil
}

func (p *pgxPoolDB) Select(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "query select")
	defer span.Finish()
	span.SetTag("query", query)

	return pgxscan.Select(ctx, p.getQuerier(ctx), dst, query, args...)
}
