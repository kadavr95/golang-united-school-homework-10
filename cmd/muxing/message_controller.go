package main

import (
	"fmt"
	"io"
	"net/http"
)

func PostMessage(w http.ResponseWriter, r *http.Request) {
	param, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = fmt.Fprintf(w, "I got message:\n"+string(param))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
