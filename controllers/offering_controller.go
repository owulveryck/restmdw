package controllers

import (
	"github.com/owulveryck/restmdw/services"
	"net/http"
)

// List all the available offerings
func OfferingListController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	// Extracting token
	status, body, err := services.OfferingList(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		next(w, r)
		return
	}
	if status != http.StatusOK {
		w.WriteHeader(status)
		next(w, r)
		return
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
		next(w, r)
		return
	}

}

// OfferingShow display the offering referenced by offeringId
func OfferingShowController(w http.ResponseWriter, r *http.Request) {
	/*
				vars := mux.Vars(r)
				var offeringId int
				var err error
				if offeringId, err = strconv.Atoi(vars["offeringId"]); err != nil {
					panic(err)
				}
					todo := RepoFindTodo(offeringId)
					if todo.Id > 0 {
						w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		yy				w.WriteHeader(http.StatusOK)
						if err := json.NewEncoder(w).Encode(todo); err != nil {
							panic(err)
						}
						return
					}
				// If we didn't find it, 404
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusNotFound)
				//if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
				//		panic(err)
				//	}
	*/

}
