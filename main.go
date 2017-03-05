package main

import (
	"database/sql"
	"fmt"
	"photo-aggregator/src/domain"

	_ "github.com/lib/pq"
)

func main() {
	ph := new(domain.Photographer)

	fmt.Println(ph.ID)

	db, err := sql.Open("postgres", "postgres://postgres:vqislemaro1@localhost/photo?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var id int
		var email string
		errr := rows.Scan(&id, &email)
		if errr != nil {
			fmt.Println(errr)
		}
		fmt.Println(id, " ", email)
	}

}
