package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["PARAM"]

	_, err := fmt.Fprintf(w, "Hello, "+param+"!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
