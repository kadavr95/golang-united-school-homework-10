package main

import (
	"net/http"
)

func GetBadUrl(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	return
}
