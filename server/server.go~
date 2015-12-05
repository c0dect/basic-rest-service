package main

import (
	"net/http"

	"gitlab.com/c0dect/basic-rest-service/routers"
)

func init() {
	router := routers.InitRoutes()
	http.Handle("/", router)
}
