package ch_5_column_type

//! MAPPING TIPE DATA
// * VARCHAR, CHAR 						= string
// * INT, BIGINT 						= int32, int64
// * FLOAT, DOUBLE 						= float32, float64
// * BOOLEAN 							= bool
// * DATE, DATETIME, TIME, TIMESTAMP 	= time.Time

import (
	"belajar-golang-database/database"
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

// type Customer struct {
// 	Id        string
// 	Name      string
// 	Email     string
// 	Balance   int32
// 	Rating    float64
// 	CreatedAt time.Time
// 	BirthDate time.Time
// 	Married   bool
// }

func TestColumnType(t *testing.T) {
	db := database.GetConnect()
	defer db.Close()

	ctx := context.Background() // membuat context
	query := "SELECT id, name, email, balance, rating, created_at, birth_date, married FROM customers"

	rows, err := db.QueryContext(ctx, query) // mequery data di databasee

	if err != nil {
		panic(err)
	}
	defer rows.Close() // kita tutup

	// * Membaca data yang ada pada rows
	for rows.Next() {
		var id, name string
		var email sql.NullString // ! Membuat tipe data string yang bisa null
		var balance int32
		var rating float64
		var created_at time.Time
		var birth_date sql.NullTime // ! membuat tipe data time yang bisa null
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}

		fmt.Println()
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		if email.Valid { // cek apakah null atau tidak
			fmt.Println("Email :", email.String)
		}
		fmt.Println("Balance :", balance)
		fmt.Println("Rating :", rating)
		fmt.Println("Created At :", created_at)
		if birth_date.Valid { // cek apakah null atau tidak
			fmt.Println("Birth Date :", birth_date.Time)
		}
		fmt.Println("Married :", married)

	}

}
