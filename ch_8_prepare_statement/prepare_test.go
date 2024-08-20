package ch_8_prepare_statement

import (
	"belajar-golang-database/database"
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestPrepareStatement(t *testing.T) {
	db := database.GetConnect()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer stmt.Close() // harus di close biar tidak terjadi memory leak

	for i := 0; i < 10; i++ {

		email := "nurdin" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komen ke " + strconv.Itoa(i)

		//* Mengeksekusi dan memasukan param
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		rowAffected, err := result.RowsAffected()
		if err != nil {
			panic(err)
		}
		lastId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Row affected :", rowAffected)
		fmt.Println("Last insert id :", lastId)
	}
}
