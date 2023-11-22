package main

import (
	"ad-server-project/src/repository"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "root", "mysql", "3306", "mysql")
	db, err := sql.Open(`mysql`, connection)
	if err != nil {
		fmt.Println("DB connection failed:", err)
		return
	}
	fmt.Println("DB connection success")

	advertisementRepo := repository.NewAdvertisementRepository(db)
	fmt.Print(advertisementRepo)
}
