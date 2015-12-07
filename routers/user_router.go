package routers

import (
	"net/http"

	"github.com/c0dect/basic-rest-service/controllers"

	"github.com/gorilla/mux"
)

var userRoutes = Routes{
	Route{
		"Registration",
		"POST",
		"/users",
		controllers.Register,
	},
	Route{
		"Login",
		"POST",
		"/users/login",
		controllers.Login,
	},
	Route{
		"Logout",
		"POST",
		"/users/logout",
		controllers.Logout,
	},
}

func SetUserRoutes(router *mux.Router) *mux.Router {
	//router = mux.NewRouter()
	for _, route := range userRoutes {
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
