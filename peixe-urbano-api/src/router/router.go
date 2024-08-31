package router

import (
	"peixe-urbano/src/router/routes"

	"github.com/gorilla/mux"
)

func Create() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
