package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gofr-dev/gofr"
)

type Car struct {
	ID     int    `json:"id"`
	Brand  string `json:"brand"`
	Model  string `json:"model"`
	Status string `json:"status"`
}

func CreateCar(c *gofr.Context, db *sql.DB) {
	var car Car
	err := c.ReadJSON(&car)
	if err != nil {
		c.JSON(http.StatusBadRequest, gofr.Map{"error": "Invalid JSON"})
		return
	}

	result, err := db.Exec("INSERT INTO cars (brand, model, status) VALUES (?, ?, ?)", car.Brand, car.Model, car.Status)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gofr.Map{"error": "Internal Server Error"})
		return
	}

	id, _ := result.LastInsertId()
	car.ID = int(id)

	c.JSON(http.StatusCreated, car)
}

func GetCars(c *gofr.Context, db *sql.DB) {
	rows, err := db.Query("SELECT * FROM cars")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gofr.Map{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var cars []Car
	for rows.Next() {
		var car Car
		err := rows.Scan(&car.ID, &car.Brand, &car.Model, &car.Status)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gofr.Map{"error": "Internal Server Error"})
			return
		}
		cars = append(cars, car)
	}

	c.JSON(http.StatusOK, cars)
}
func UpdateCar(c *gofr.Context, db *sql.DB) {
	var car Car
	err := c.ReadJSON(&car)
	if err != nil {
		c.JSON(http.StatusBadRequest, gofr.Map{"error": "Invalid JSON"})
		return
	}

	_, err = db.Exec("UPDATE cars SET brand=?, model=?, status=? WHERE id=?", car.Brand, car.Model, car.Status, car.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gofr.Map{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, car)
}

func DeleteCar(c *gofr.Context, db *sql.DB) {
	id, err := c.PathParamInt("id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gofr.Map{"error": "Invalid car ID"})
		return
	}

	_, err = db.Exec("DELETE FROM cars WHERE id=?", id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gofr.Map{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gofr.Map{"message": "Car deleted successfully"})
}
