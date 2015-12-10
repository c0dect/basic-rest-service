package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc *negroni.Negroni
}

type Routes []Route

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetProductRoutes(router)
	router = SetUserRoutes(router)
	return router
}
