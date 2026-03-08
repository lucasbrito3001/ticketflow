package out

import (
	"context"
	"database/sql"

	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/app/ports"
)

type MySQLUnitOfWork struct {
	db *sql.DB
}

func NewMySQLUnitOfWork(db *sql.DB) ports.UnitOfWork {
	return &MySQLUnitOfWork{
		db: db,
	}
}

func (u *MySQLUnitOfWork) Do(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	txCtx := context.WithValue(ctx, "tx", tx)
	if err := fn(txCtx); err != nil {
		return err
	}

	return tx.Commit()
}

type QueryExecutor interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

func GetQueryExecutor(ctx context.Context, db *sql.DB) QueryExecutor {
	if tx, ok := ctx.Value("tx").(*sql.Tx); ok {
		return tx
	}
	return db
}
