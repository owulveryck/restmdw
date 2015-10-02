package controllers

import (
	//"encoding/json"
	//"io"
	//"io/ioutil"
	"net/http"
	//"strconv"

	//"github.com/gorilla/mux"
)

// List all the available requests
func RequestsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//if err := json.NewEncoder(w).Encode(todos); err != nil {
	//	panic(err)
	//}
}

// RequestShow display the request referenced by requestId
func RequestShow(w http.ResponseWriter, r *http.Request) {
	/*
		vars := mux.Vars(r)
		var requestId int
		var err error
		if requestId, err = strconv.Atoi(vars["requestId"]); err != nil {
			panic(err)
		}
			todo := RepoFindTodo(requestId)
			if todo.Id > 0 {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusOK)
				if err := json.NewEncoder(w).Encode(todo); err != nil {
					panic(err)
				}
				return
			}

		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	*/

}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos

*/
func RequestCreate(w http.ResponseWriter, r *http.Request) {
	/*
		var todo Todo
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}
		if err := json.Unmarshal(body, &todo); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}

		t := RepoCreateTodo(todo)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(t); err != nil {
			panic(err)
		}
	*/
}
