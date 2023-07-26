package db

import (
	"fmt"
	"golang-graphql-gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var DB *gorm.DB

func InitializeDb(dbConfig DbConfig){
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	dbConfig.DBHost,
	dbConfig.DBPort,
	dbConfig.DBUser,
	dbConfig.DBPassword,
	dbConfig.DBName,
)
db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	checkErr(err)

	fmt.Println("konek ke pg")
	
	db.AutoMigrate(&models.Car{})


	DB = db
}