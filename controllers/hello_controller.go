package controllers

import (
	"github.com/owulveryck/restmdw/web"
	"log"
	"net/http"
)

// @Title HelloController
// @Description Says hello
// @Accept  json
// @Param   customer_id     path    int     true        "Customer ID"
// @Param   order_id        query   int     false        "Retrieve order with given ID only"
// @Param   order_nr        query   string  false        "Retrieve order with given number only"
// @Param   created_from    query   string  false        "Date-time string, MySQL format. If specified, API will retrieve orders that were created starting from created_from"
// @Param   created_to      query   string  false        "Date-time string, MySQL format. If specified, API will retrieve orders that were created before created_to"
// @Success 200
// @Failure 400 {object} web.JsonErr    "Customer ID must be specified"
// @Resource /hello
// @Router /hello [get]
func HelloController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var jsonErr web.JsonErr
	log.Println(jsonErr)
	w.Write([]byte("Hello, World!"))

}
