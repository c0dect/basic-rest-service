package main

import (
	"net/http"

	"github.com/c0dect/basic-rest-service/routers"
)

func init() {
	router := routers.InitRoutes()
	http.Handle("/", router)
}

func main() {

}
