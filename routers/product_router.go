package routers

import (
	"github.com/c0dect/basic-rest-service/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

var routes = Routes{
	Route{
		"Default",
		"GET",
		"/",
		negroni.New(
			negroni.HandlerFunc(controllers.AuthenticateUser),
			negroni.HandlerFunc(controllers.Index),
		),
	},
	Route{
		"ProductIndex",
		"GET",
		"/products",
		negroni.New(
			negroni.HandlerFunc(controllers.AuthenticateUser),
			negroni.HandlerFunc(controllers.GetProducts),
		),
	},
	Route{
		"ProductDetails",
		"GET",
		"/products/{productId}",
		negroni.New(
			negroni.HandlerFunc(controllers.AuthenticateUser),
			negroni.HandlerFunc(controllers.GetProduct),
		),
	},
	Route{
		"CreateProduct",
		"POST",
		"/products",
		negroni.New(
			negroni.HandlerFunc(controllers.AuthenticateUser),
			negroni.HandlerFunc(controllers.CreateProduct),
		),
	},
	Route{
		"ProductDelete",
		"DELETE",
		"/products/{productId}",
		negroni.New(
			negroni.HandlerFunc(controllers.AuthenticateUser),
			negroni.HandlerFunc(controllers.DeleteProduct),
		),
	},
	Route{
		"ProductUpdate",
		"PATCH",
		"/products/{productId}",
		negroni.New(
			negroni.HandlerFunc(controllers.AuthenticateUser),
			negroni.HandlerFunc(controllers.UpdateProduct),
		),
	},
}

func SetProductRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {

		router.
			Handle(route.Pattern, route.HandlerFunc).
			Methods(route.Method).
			Name(route.Name)
	}
	return router
}
