package routers

import (
	"github.com/gorilla/mux"
	"github.com/moehandi/campaign/api/controllers"
)

func SetAuthRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/auth/register", controllers.AuthRegister).Methods("POST")
	router.HandleFunc("/auth/token", controllers.AuthToken).Methods("POST")
	return router
}
