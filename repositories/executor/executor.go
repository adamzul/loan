package executor

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"loan.com/connection"
)

type Executor struct {
	db *connection.ReplicationDB
}

func New(db *connection.ReplicationDB) Executor {
	return Executor{db}
}

type GetterContext interface {
	GetContext(ctx context.Context, dest any, query string, args ...any) error
}

type SelectorContext interface {
	SelectContext(ctx context.Context, dest any, query string, args ...any) error
}

func (e *Executor) GetContext(ctx context.Context, dest any, query string, args ...any) error {
	var getter GetterContext

	ok, tx := IsTransaction(ctx)
	if ok {
		getter = tx
	} else {
		getter = e.db.Standby
	}

	return getter.GetContext(ctx, dest, query, args...)
}

func (e *Executor) SelectContext(ctx context.Context, dest any, query string, args ...any) error {
	var selector SelectorContext

	ok, tx := IsTransaction(ctx)
	if ok {
		selector = tx
	} else {
		selector = e.db.Standby
	}

	return selector.SelectContext(ctx, dest, query, args...)
}

func (e *Executor) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	ok, tx := IsTransaction(ctx)
	if ok {
		return tx.ExecContext(ctx, query, args...)
	}

	return e.db.Primary.ExecContext(ctx, query, args...)
}

func (e *Executor) QueryRowxScanContext(ctx context.Context, dest any, query string, args ...any) error {
	ok, tx := IsTransaction(ctx)
	if ok {
		return tx.QueryRowxContext(ctx, query, args...).StructScan(dest)
	}
	return e.db.Primary.QueryRowxContext(ctx, query, args...).StructScan(dest)
}

func (e *Executor) QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error) {
	ok, tx := IsTransaction(ctx)
	if ok {
		return tx.QueryxContext(ctx, query, args...)
	}

	return e.db.Primary.QueryxContext(ctx, query, args...)
}

func (e *Executor) NamedExecContext(ctx context.Context, query string, arg any) (sql.Result, error) {
	ok, tx := IsTransaction(ctx)
	if ok {
		return tx.NamedExecContext(ctx, query, arg)
	}

	return e.db.Primary.NamedExecContext(ctx, query, arg)
}
