package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func RoutesMap(api *ApiConnection) Routes {
	var routes = Routes{
		Route{"Index", "GET", "/", api.Index},
		Route{"SignIn", "POST", "/signin", api.uHandlers.SignIn},
		Route{"Login", "POST", "/login", api.uHandlers.Login},
	}

	return routes
}
