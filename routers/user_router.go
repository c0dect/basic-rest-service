package routers

import (
	"github.com/c0dect/basic-rest-service/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

var userRoutes = Routes{
	Route{
		"Registration",
		"POST",
		"/users",
		negroni.New(
			negroni.HandlerFunc(controllers.Register),
		),
	},
	Route{
		"Login",
		"POST",
		"/users/login",
		negroni.New(
			negroni.HandlerFunc(controllers.Login),
		),
	},
	Route{
		"Logout",
		"POST",
		"/users/logout",
		negroni.New(
			negroni.HandlerFunc(controllers.AuthenticateUser),
			negroni.HandlerFunc(controllers.Logout),
		),
	},
}

func SetUserRoutes(router *mux.Router) *mux.Router {
	for _, route := range userRoutes {
		router.
			Handle(route.Pattern, route.HandlerFunc).
			Methods(route.Method).
			Name(route.Name)
	}
	return router
}
