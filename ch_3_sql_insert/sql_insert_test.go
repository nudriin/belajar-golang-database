package ch_3_sql_insert

import (
	"belajar-golang-database/database"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsertData(t *testing.T) {
	db := database.GetConnect()
	defer db.Close()

	ctx := context.Background() // membuat contexnya

	query := "INSERT INTO customers(id, name) VALUES('C001', 'Nurdin')"
	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert data")

}
