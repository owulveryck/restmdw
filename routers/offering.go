package routers

import (
	"github.com/owulveryck/restmdw/controllers"
	"github.com/owulveryck/restmdw/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetOfferingRoutes(router *mux.Router) *mux.Router {
	router.Handle("/offering",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.OfferingListController),
		)).Methods("GET")

	return router
}
