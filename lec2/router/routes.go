package router

import (
	"net/http"

	"btap/golang/controller"
)

func CalculateRouter() {
	http.HandleFunc("/", controller.Handler)
}