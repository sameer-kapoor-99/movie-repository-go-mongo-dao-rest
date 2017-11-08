package moviestore

import (
	"net/http"
	"github.com/gorilla/mux"
	"../dao/interfaces"
)



// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route



//NewRouter configures a new router to the API
func NewRouter(dao interfaces.UserDao) *mux.Router {
	var controller = &Controller{dao : dao}

	var routes = Routes{
		Route{
			"Index",
			"GET",
			"/movies/",
			controller.Index,
		},
		Route{
			"AddMovie",
			"POST",
			"/movies/",
			controller.AddMovie,
		},
		Route{
			"UpdateMovie",
			"PUT",
			"/movies/{id}",
			controller.UpdateMovie,
		},
		Route{
			"DeleteMovie",
			"DELETE",
			"/movies/{id}",
			controller.DeleteMovie,
		},
		Route{
			"GetMovieByTitle",
			"GET",
			"/movies/title/{title}",
			controller.GetMovieByTitle,
		},
	}

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
