package main

import (
	"github.com/gofr-dev/gofr"
)

func initDB() {

}

func CreateCar(ctx *gofr.Context) {

}

func GetCars(ctx *gofr.Context) {

}

func UpdateCar(ctx *gofr.Context) {

}

func DeleteCar(ctx *gofr.Context) {

}

func main() {
	initDB()

	app := gofr.New()
	app.POST("/cars", CreateCar)
	app.GET("/cars", GetCars)
	app.PUT("/cars", UpdateCar)
	app.DELETE("/cars", DeleteCar)

	app.Run(":8080")
}
