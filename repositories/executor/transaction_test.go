package executor

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"loan.com/connection"
)

var (
	errSQLMock = errors.New("sql: mock error")
)

func TestTransaction_OK(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("repo.postgesql.Transaction: an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	parentCtx := context.Background()
	store := &connection.ReplicationDB{
		Primary: sqlx.NewDb(db, "sqlmock"),
	}
	trx := NewTransaction(store)
	err = trx.Execute(parentCtx, func(ctx context.Context) error {
		ok, tx := IsTransaction(ctx)
		if !ok {
			t.Fatalf("repo.postgesql.Transaction: invalid context wrapper")
		}

		_, errTrx := tx.Exec("UPDATE products SET views = views + 1")
		return errTrx
	})

	if err != nil {
		t.Fatalf("repo.postgesql.Transaction: error transaction %s", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("repo.postgesql.Transaction: error doesn't meet expectations %s", err)
	}
}

func TestTransaction_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("repo.postgesql.Transaction: an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectRollback()

	parentCtx := context.Background()
	store := &connection.ReplicationDB{
		Primary: sqlx.NewDb(db, "sqlmock"),
	}

	trx := NewTransaction(store)
	err = trx.Execute(parentCtx, func(ctx context.Context) error {
		ok, tx := IsTransaction(ctx)
		if !ok {
			t.Fatalf("repo.postgesql.Transaction: invalid context wrapper")
		}

		_, _ = tx.Exec("UPDATE products SET views = views + 1")
		return errSQLMock
	})

	if err != errSQLMock {
		t.Fatalf("repo.postgesql.Transaction: error transaction %s", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("repo.postgesql.Transaction: error doesn't meet expectations %s", err)
	}
}

func TestTransaction_Panic(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("repo.postgesql.Transaction: an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectRollback()

	parentCtx := context.Background()
	store := &connection.ReplicationDB{
		Primary: sqlx.NewDb(db, "sqlmock"),
	}

	trx := NewTransaction(store)
	err = trx.Execute(parentCtx, func(ctx context.Context) error {
		ok, tx := IsTransaction(ctx)
		if !ok {
			t.Fatalf("repo.postgesql.Transaction: invalid context wrapper")
		}

		_, _ = tx.Exec("UPDATE products SET views = views + 1")
		panic(errSQLMock)
	})

	if err != errSQLMock {
		t.Fatalf("repo.postgesql.Transaction: error transaction %s", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("repo.postgesql.Transaction: error doesn't meet expectations %s", err)
	}
}
