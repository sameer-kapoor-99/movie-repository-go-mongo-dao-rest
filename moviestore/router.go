package moviestore

import (
	"net/http"
	"github.com/gorilla/mux"
)

var controller = &Controller{}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"AddMovie",
		"POST",
		"/addmovie/{id}",
		controller.AddMovie,
	},
	Route{
		"UpdateMovie",
		"PUT",
		"/",
		controller.UpdateMovie,
	},
	Route{
		"DeleteMovie",
		"DELETE",
		"/delete/{id}",
		controller.DeleteMovie,
	},
	Route{
		"GetMovieByTitle",
		"GET",
		"/title/{title}",
		controller.GetMovieByTitle,
	},
}

//NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
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
