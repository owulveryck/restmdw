package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetHelloRoutes(router)
	router = SetOfferingRoutes(router)
	router = SetAuthenticationRoutes(router)
	return router
}
