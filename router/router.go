package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/padwalab/goservices/testsvc"
)

// Route to hold the structural design of each route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a list of endpoints microservice is going to expose
type Routes []Route

var routes = Routes{
	Route{
		"Test",
		"GET",
		"/api/test",
		testsvc.SampleService,
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
