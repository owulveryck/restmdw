package main

import (
	"github.com/owulveryck/restmdw/routers"
	"github.com/owulveryck/restmdw/settings"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}
