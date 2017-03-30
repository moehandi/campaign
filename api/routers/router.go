package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	router = SetAuthRoutes(router)

	router = SetSurveyRoutes(router)

	//router = SetUserRoutes(router)

	return router
}
