package ch_7_last_insert_id

import (
	"belajar-golang-database/database"
	"context"
	"fmt"
	"testing"
)

func TestLastInsertId(t *testing.T) {
	db := database.GetConnect()
	defer db.Close()

	email := "samehada@gmail.com"
	comment := "Keep it cool"
	ctx := context.Background()
	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId() // mengambil last insert id
	if err != nil {
		panic(err)
	}
	fmt.Println("Last id :", insertId)

	rowAffected, err := result.RowsAffected() // mengambil jumlah row yang terkena effect
	if err != nil {
		panic(err)
	}
	fmt.Println("Row affected :", rowAffected)

}
