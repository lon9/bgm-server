package utils

import "net/http"

//CheckAuth is check basic auth.
func CheckAuth(r *http.Request) bool {
	username, password, ok := r.BasicAuth()

	if ok == false {
		return false
	}
	return username == "BGMAdmin" && password == "admin"
}
