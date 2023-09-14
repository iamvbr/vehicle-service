package main

import (
	"github.com/iamvbr/vehicle-service/internal/http/cars"
	"github.com/iamvbr/vehicle-service/internal/middlewares"
	cars3 "github.com/iamvbr/vehicle-service/internal/services/cars"
	cars2 "github.com/iamvbr/vehicle-service/internal/stores/cars"
	"net/http"
)

func main() {
	carStore := cars2.New()
	carSVC := cars3.New(carStore)
	carsHandler := cars.New(carSVC)

	http.ListenAndServe(":8080", middlewares.Logger(carsHandler))
}
