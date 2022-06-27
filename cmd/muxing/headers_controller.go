package main

import (
	"net/http"
	"strconv"
	"strings"
)

func PostHeaders(w http.ResponseWriter, r *http.Request) {
	var responseHeaders []string
	responseResult := 0

	for k, v := range r.Header {
		// Most likely it was enough to just hardcode A and B headers ¯\_(ツ)_/¯
		if len(k) <= 2 {
			vInt, err := strconv.Atoi(strings.Join(v, ""))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			responseHeaders = append(responseHeaders, k)
			responseResult += vInt
		}
	}

	w.Header().Set(strings.Join(responseHeaders, "+"), strconv.Itoa(responseResult))

	w.WriteHeader(http.StatusOK)
	return
}
