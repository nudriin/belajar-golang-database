package ch_1_connection

import (
	"database/sql"
	"testing"

	// * Memanggil init function dari drivers
	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:hishasyl115a1408@tcp(localhost:3306)/belajar_go_db")

	if err != nil {
		panic(err)
	}

	defer db.Close() // menutup database agar tidak terjadi leak
}
