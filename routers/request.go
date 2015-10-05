package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/owulveryck/restmdw/controllers"
	"github.com/owulveryck/restmdw/core/authentication"
)

func SetRequestRoutes(router *mux.Router) *mux.Router {
	router.Handle("/request",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.RequestListController),
		)).Methods("GET")
	router.Handle("/request",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.RequestShowController),
		)).Methods("POST")

	return router
}
