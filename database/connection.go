package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnect() *sql.DB {
	// parseTime=true //* agar bisa konversi tipe data date atau timestamp atau datetime dll
	db, err := sql.Open("mysql", "root:hishasyl115a1408@tcp(localhost:3306)/belajar_go_db?parseTime=true")

	if err != nil {
		panic(err)
	}

	// * Setting connection pooling
	db.SetMaxIdleConns(10)                  // membuat minimal koneksi yang tersedia
	db.SetMaxOpenConns(100)                 // membuat maksimal koneksi yang dapat dibuka
	db.SetConnMaxIdleTime(5 * time.Minute)  // membuat maksimal waktu idle koneksi, apabila da koneksi terbuka dan tidak digukan maka akan di close
	db.SetConnMaxLifetime(60 * time.Minute) // setelah 1 jam koneksi akan diperbaharui

	return db
}
