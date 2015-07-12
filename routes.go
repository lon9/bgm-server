package main

import "net/http"

// Route is struct for router.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is list of route.
type Routes []Route

var routes = Routes{
	Route{
		"Proxy",
		"GET",
		"/proxy",
		proxy,
	},
}
