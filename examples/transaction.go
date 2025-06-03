package examples

import (
	"context"
	"database/sql"
	"fmt"
	ctxtx "saifutdinov/ca-ctx-tx"
	"saifutdinov/ca-ctx-tx/examples/repository/sqldb"
	"saifutdinov/ca-ctx-tx/examples/usecase"
)

func Init() {

	// Create a new database connection
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Create a new dbs instance
	repo := sqldb.NewRepository(db)

	// Create a new txs instance
	txs := ctxtx.NewTransaction(db)

	// Create a new usecase instance
	usecase := usecase.NewUsecase(txs, repo)

	// your context
	ctx := context.Background()

	usecase.DoSomething(ctx)
}
