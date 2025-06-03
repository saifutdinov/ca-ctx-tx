package sqldb

import (
	"context"
	"database/sql"
	ctxtx "saifutdinov/ca-ctx-tx"
	"saifutdinov/ca-ctx-tx/examples/domain"
)

type Repository struct {
	ctxtx.DBI
}

func NewRepository(db *sql.DB) domain.Repository {
	return &Repository{
		// now default db has out overlay with our methods with transactions.
		DBI: ctxtx.NewDBS(db),
	}
}

func (r *Repository) DoSomething(ctx context.Context) {
	// instide of each mehod we check ctx for exists transaction.
	// if we found sql.Tx, current method running in this transaction.
	// Othervise we run query in default db. (ctxtx.NewDBS(db))
	r.ExecContext(ctx, `INSERT INTO foo (id) VALUES (1)`)
}

func (r *Repository) DoSomethingElse(ctx context.Context) {
	r.ExecContext(ctx, `INSERT INTO foo (id) VALUES (2)`)
}

func (r *Repository) DoSomethingAgain(ctx context.Context) {
	r.QueryContext(ctx, `SELECT * FROM foo`)
}

func (r *Repository) DoSomethingElseAgain(ctx context.Context) {
	r.QueryRowContext(ctx, `SELECT * FROM foo WHERE id = 1`)
}

func (r *Repository) DoSomethingElseAgainAgain(ctx context.Context) {
	r.ExecContext(ctx, `UPDATE foo SET id = 2 WHERE id = 1`)
}
