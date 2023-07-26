package controllers

import (
	"golang-graphql-gorm/db"
	"golang-graphql-gorm/models"

	"gorm.io/gorm"
)

func GetCar(id int)([]models.Car, error){
	car := new([]models.Car)

	var result *gorm.DB
	if id != 0 {
		result = db.DB.Debug().Where(id).Find(car)
	} else{
		result = db.DB.Find(car)
	}
	return *car, result.Error

}

// func DeleteCar(id int) (c *gin.Context){
func DeleteCar(id int) ([]models.Car, error){

	var car models.Car
	db.DB.Debug().Where("id = ?", id).Find(&car)
	db.DB.Debug().Delete(&car)
	
	// c.JSON(http.StatusOK, gin.H{"data": car})
	return nil, nil
	// return nil, nil
}

// func CreateCar(carData{}){
// 	var car models.Car
// 	db.DB.Debug().Create(&car)
// }