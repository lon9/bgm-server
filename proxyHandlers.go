package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func proxy(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	url := params.Get("url")

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Fprintln(w, string(body))

}
