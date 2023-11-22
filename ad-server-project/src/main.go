package main

import (
	"ad-server-project/src/adapter/http"
	"ad-server-project/src/repository"
	"ad-server-project/src/usecase"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	ht "net/http"
)

func main() {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "root", "mysql", "3306", "ad_server_project")
	db, err := sql.Open(`mysql`, connection)
	if err != nil {
		fmt.Println("DB connection failed:", err)
		return
	}
	fmt.Println("DB connection success")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(ht.StatusOK, "Hello, World!")
	})

	advertisementRepo := repository.NewAdvertisementRepository(db)
	advertisementUsecase := usecase.NewAdvertisementUsecase(advertisementRepo)
	http.NewAdvertisementHandler(router, advertisementUsecase)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}

}
