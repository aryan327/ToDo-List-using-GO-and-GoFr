package routes

import (
	"https://github.com/gofr-dev/gofr"
	"github.com/ichtrojan/go-todo/controllers"
)

func Init() *gofr.Router {
	route := gofr.NewRouter()

	route.HandleFunc("/", controllers.Show)
	route.HandleFunc("/add", controllers.Add).Methods("POST")
	route.HandleFunc("/delete/{id}", controllers.Delete)
	route.HandleFunc("/complete/{id}", controllers.Complete)

	return route
}
