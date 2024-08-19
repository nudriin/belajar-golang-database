package ch_4_sql_query

import (
	"belajar-golang-database/database"
	"context"
	"fmt"
	"testing"
)

func TestQueryData(t *testing.T) {
	db := database.GetConnect()
	defer db.Close()

	ctx := context.Background() // membuat context
	query := "SELECT * FROM customers"

	rows, err := db.QueryContext(ctx, query) // mequery data di databasee

	if err != nil {
		panic(err)
	}
	defer rows.Close() // kita tutup

	// * Membaca data yang ada pada rows
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id :", id)
		fmt.Println("Name :", name)

	}

}
