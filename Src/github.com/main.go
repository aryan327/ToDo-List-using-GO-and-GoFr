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

func main() {
	initDB()

	app := gofr.New()
	app.POST("/cars", CreateCar)
	app.GET("/cars", GetCars)


	app.Run(":8080")
}
