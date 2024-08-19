package ch_6_sql_with_param

import (
	"belajar-golang-database/database"
	"context"
	"fmt"
	"testing"
)

func TestQueryWithParam(t *testing.T) {
	db := database.GetConnect()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"
	query := "SELECT username FROM users WHERE username = ? AND password = ? LIMIT 1" // * melalukan query param

	rows, err := db.QueryContext(ctx, query, username, password) // * menginputkan parameter

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success login with username =", username)
	} else {
		fmt.Println("Username or password is wrong")
	}
}

func TestExecWithParam(t *testing.T) {
	db := database.GetConnect()
	defer db.Close()

	ctx := context.Background()

	username := "nurdin"
	password := "nurdin"
	query := "INSERT INTO users(username, password) VALUES(?, ?)" // * melalukan exec param

	_, err := db.ExecContext(ctx, query, username, password) // * menginputkan parameter

	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses input data")

}
