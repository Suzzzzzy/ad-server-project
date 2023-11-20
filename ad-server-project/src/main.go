package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "root", "localhost", "3306", "mysql")
	db, err := sql.Open(`mysql`, connection)
	fmt.Println("DB connection Success")

	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

}
