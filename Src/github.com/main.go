package main

import (
	"github.com/gofr-dev/gofr"
)

func initDB() {
dsn := "user:password@tcp(localhost:8080)/dbname?parseTime=true"
var err error
db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
if err != nil {
	panic("Failed to connect to the database: " + err.Error())
}

func CreateCar(ctx *gofr.Context) {
	var newCar Car
	err := ctx.JSONBody(&newCar)
	if err != nil {
		// Handle the error, e.g., invalid JSON
		ctx.JSON(400, map[string]interface{}{"error": "Invalid JSON"})
		return
	}

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
