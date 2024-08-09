package executor

import (
	"context"

	"github.com/jmoiron/sqlx"
	"loan.com/connection"
)

type ContextTxKey struct{}

type ContextTxWrapper struct {
	tx *sqlx.Tx
}

type Transaction interface {
	Execute(ctx context.Context, fn func(ctx context.Context) error) error
}

type transaction struct {
	db *connection.ReplicationDB
}

func NewTransaction(db *connection.ReplicationDB) Transaction {
	return &transaction{
		db: db,
	}
}

func (t *transaction) Execute(ctx context.Context, fn func(ctx context.Context) error) error {
	// use primary for transaction
	tx, err := t.db.Primary.Beginx()
	if err != nil {
		return err
	}

	ctxTx := context.WithValue(ctx, ContextTxKey{}, ContextTxWrapper{tx: tx})

	func() {
		defer func() {
			if p := recover(); p != nil {
				// keep original error
				_ = tx.Rollback()
				switch e := p.(type) {
				case error:
					err = e
				default:
					panic(e)
				}
			} else if err != nil {
				_ = tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()
		err = fn(ctxTx)
	}()

	return err
}

// IsTransaction checks wether context contain transaction or not
func IsTransaction(ctx context.Context) (bool, *sqlx.Tx) {
	ctxTxWrapper, ok := ctx.Value(ContextTxKey{}).(ContextTxWrapper)

	return ok, ctxTxWrapper.tx
}
