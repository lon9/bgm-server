package main

import (
	"net/http"
)

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
	Route{
		"InqueryCreate",
		"POST",
		"/inquery",
		inqueryCreate,
	},
	Route{
		"InqueryIndex",
		"GET",
		"/inquery",
		inqueryIndex,
	},
	Route{
		"InqueryShow",
		"GET",
		"/inquery/{inqueryId}",
		inqueryShow,
	},
	Route{
		"InqueryDelete",
		"DELETE",
		"/inquery/{inqueryId}",
		inqueryDelete,
	},
	Route{
		"VideoIndex",
		"GET",
		"/video",
		videoIndex,
	},
	Route{
		"VideoUodate",
		"POST",
		"/video/{videoId}",
		videoUpdate,
	},
	Route{
		"LikeUpdate",
		"POST",
		"/like",
		likeUpdate,
	},
}
