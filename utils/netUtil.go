package utils

import (
	"fmt"
	"net/http"
)

// CheckError is inspect error and read it.
func CheckError(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Fprintln(w, err)
	}
}

// SetJSONHeader is add header to json response.
func SetJSONHeader(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return w
}
