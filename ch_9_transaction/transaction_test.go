package ch_9_transaction

import (
	"belajar-golang-database/database"
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestTransaction(t *testing.T) {
	db := database.GetConnect()
	defer db.Close()

	ctx := context.Background()

	// * memulai transaction
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	for i := 0; i < 10; i++ {
		email := "nurdin" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komen ke " + strconv.Itoa(i)

		//* Mengeksekusi dan memasukan param
		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment id ke", id)

	}

	// * Meng commit transaction
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
