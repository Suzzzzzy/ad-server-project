package main

import (
	"ad-server-project/src/domain/model"
	"ad-server-project/src/repository"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "root", "localhost", "3306", "mysql")
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		fmt.Println("DB connection failed:", err)
		return
	}
	fmt.Println("DB connection success")

	// Auto Migrate: 데이터베이스 테이블 자동 생성
	err = db.AutoMigrate(&model.User{}, &model.Advertisement{})
	if err != nil {
		fmt.Println("Auto Migrate failed:", err)
		return
	}

	advertisementRepo := repository.NewAdvertisementRepository(db)

}
