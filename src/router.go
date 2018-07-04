package src

import (
	"github.com/gorilla/mux"
	"log"
	"./controller"
	"./repository"
	"net/http"
)

var c = &controller.Controller{Repository:repository.ListItemRepository{}}

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes = []Route
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		c.Index,
	},
	Route{
		"Insert",
		"POST",
		"/",
		c.Insert,
	},
}

func NewRouter() *mux.Router{
	router:= mux.NewRouter().StrictSlash(true)
	for _, route := range routes{
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	return router
}
