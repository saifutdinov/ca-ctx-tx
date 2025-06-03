package examples

import (
	"context"
	"database/sql"
	"fmt"
	ctxtx "saifutdinov/ca-ctx-tx"
)

func SimpleQuery() {
	// Create a new database connection
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Create a new dbs instance
	dbs := ctxtx.NewDBS(db)

	// your context
	ctx := context.Background()

	// Execute a query
	rows, err := dbs.QueryContext(ctx, "SELECT * FROM table")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	// Print the results
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
}
