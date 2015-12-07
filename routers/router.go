package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetProductRoutes(router)
	router = SetUserRoutes(router)
	return router
}
