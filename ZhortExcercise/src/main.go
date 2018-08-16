package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	func1 := func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		w.Write([]byte(vars["id"]))

	}

	r.HandleFunc("/foo/{id}", func1)

	http.Handle("/", r)
	go http.ListenAndServe(":4000", nil)
	fmt.Scanln()
}
