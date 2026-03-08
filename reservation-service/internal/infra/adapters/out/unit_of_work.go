package out

import (
	"context"
	"database/sql"

	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/ports"
)

type (
	txKey struct{}

	UnitOfWork struct {
		db *sql.DB
	}

	QueryExecutor interface {
		ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
		QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
		QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	}
)

func NewUnitOfWork(db *sql.DB) ports.UnitOfWork {
	return &UnitOfWork{db: db}
}

func (u *UnitOfWork) Do(
	ctx context.Context,
	fn func(ctx context.Context) error,
) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	txCtx := context.WithValue(ctx, txKey{}, tx)

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}

		if err != nil {
			_ = tx.Rollback()
		}
	}()

	err = fn(txCtx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

func GetQueryExecutor(ctx context.Context, db *sql.DB) QueryExecutor {
	if tx := extractTxFromContext(ctx); tx != nil {
		return tx
	}
	return db
}

func extractTxFromContext(ctx context.Context) *sql.Tx {
	tx, ok := ctx.Value(txKey{}).(*sql.Tx)
	if !ok {
		return nil
	}

	return tx
}
