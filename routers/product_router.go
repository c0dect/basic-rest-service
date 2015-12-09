package routers

import (
	"net/http"

	"github.com/c0dect/basic-rest-service/controllers"

	"github.com/gorilla/mux"
)

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
		controllers.AuthenticateUser(controllers.CreateProduct),
	},
}

func SetProductRoutes(router *mux.Router) *mux.Router {
	//router = mux.NewRouter()
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
