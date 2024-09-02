package routes

import (
	"net/http"
	"peixe-urbano/src/controllers"
)

var routesOffers = []Route{
	{
		URI:                    "/offers",
		Method:                 http.MethodPost,
		Function:               controllers.CreateOffer,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/offers",
		Method:                 http.MethodGet,
		Function:               controllers.GetOffers,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/offers/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.GetOffer,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/offers/{userId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateOffer,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/offers/{userId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteOffer,
		RequiredAuthentication: false,
	},
}
