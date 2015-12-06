package routers

import (
	"net/http"

	"github.com/c0dect/basic-rest-service/controllers"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Default",
		"GET",
		"/",
		controllers.Index,
	},
	Route{
		"ProductIndex",
		"GET",
		"/products",
		controllers.GetProducts,
	},
	Route{
		"CreateProduct",
		"POST",
		"/products",
		controllers.CreateProduct,
	},
}

func SetProductRoutes(router *mux.Router) *mux.Router {
	router = mux.NewRouter()
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
