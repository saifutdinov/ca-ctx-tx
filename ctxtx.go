// The ctxtx package provides a set of functions and types for working with a database,
// including executing queries, preparing statements, and querying rows.
//
// The methods are designed to be used with a context.Context
// object, which is used to manage the execution of the database operations.
//
// The package is designed to be used with a variety of database drivers, and provides
// a common interface for working with different databases.
//
// Main function is switchDB. If current context has value
// (and we know that sqlTxValueKey has exact this value is context.Value)
// then it returns sql.Tx that means this query
// (one of the ExecContext, QueryContext, QueryRowContext)
// is executed in transaction.
//
// Otherwise it returns sql.DB and this query is executed in database directly.
package ctxtx

// ctxtx.go provides TXS interface to begin,commit or rollback transaction in your usecase(business-logic).

import (
	"context"
	"database/sql"
)

// sqlTxValueKey key for sql transaction stored in child context.
const sqlTxValueKey = sqlTxKeyType("sql_tx_value_key")

type (
	// sqlTxKeyType type for sql transaction context key
	sqlTxKeyType string

	// TXI interface for BeginTx, CommitTx, or RollbackTx
	TXS struct {
		dbptr *sql.DB
	}
)

func NewTransaction(dbp *sql.DB) TXI {
	return &TXS{dbptr: dbp}
}

func switchDB(ctx context.Context, dbp *sql.DB) DBI {
	if tx := ctx.Value(sqlTxValueKey); tx != nil {
		return tx.(*sql.Tx)
	} else {
		return dbp
	}
}

func (txs *TXS) BeginTx(ctx context.Context) (context.Context, error) {
	tx, err := txs.dbptr.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, sqlTxValueKey, tx), nil
}

func (txs *TXS) CommitTx(ctx context.Context) error {
	if tx := ctx.Value(sqlTxValueKey); tx != nil {
		if err := tx.(*sql.Tx).Commit(); err != nil && err != sql.ErrTxDone {
			// fmt.Println("Transaction, CommitTx: ", err)
			return err
		}
	}
	return nil
}

func (txs *TXS) RollbackTx(ctx context.Context) error {
	if tx := ctx.Value(sqlTxValueKey); tx != nil {
		if err := tx.(*sql.Tx).Rollback(); err != nil && err != sql.ErrTxDone {
			// fmt.Println("Transaction, RollbackTx: ", err)
			return err
		}
	}
	return nil
}
