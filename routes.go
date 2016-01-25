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
		Route{"UpdateProfile", "POST", "/update_profile", api.uHandlers.UpdateProfile},
		Route{"ReadProfile", "POST", "/read_profile", api.uHandlers.ReadUserProfile},
		Route{"CreateStore", "POST", "/create_store", api.sHandlers.CreateStore},
		Route{"UpdateStore", "POST", "/update_store", api.sHandlers.UpdateStore},
		Route{"DeleteStore", "POST", "/update_store", api.sHandlers.DeleteStore},
		Route{"ListPersonalStores", "POST", "/personal_stores", api.sHandlers.ListPersonalStores},
		Route{"ListStores", "POST", "/list_stores", api.sHandlers.ListStores},
		Route{"CreateProduct", "POST", "/create_product", api.pHandlers.CreateProduct},
		Route{"UpdateProduct", "POST", "/update_product", api.pHandlers.UpdateProduct},
		Route{"DeleteProduct", "POST", "/delete_product", api.pHandlers.DeleteProduct},
		Route{"ListProducts", "POST", "/list_product", api.pHandlers.ListProductsByStore},
	}

	return routes
}
