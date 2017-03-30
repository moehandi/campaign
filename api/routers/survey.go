package routers

import (
	"github.com/gorilla/mux"
	"github.com/moehandi/gokyl"
	//"github.com/moehandi/api/common"
	"github.com/moehandi/campaign/api/controllers"
)

const (
	v1 = "/api/v1"
)
func SetSurveyRoutes(router *mux.Router) *mux.Router {
	surveyRouter := mux.NewRouter()
	surveyRouter.StrictSlash(true)
	surveyRouter.HandleFunc(v1 + "/surveys/", controllers.CreateSurvey).Methods("POST")
	surveyRouter.HandleFunc(v1 + "/surveys/", controllers.GetSurveys).Methods("GET")
	surveyRouter.HandleFunc(v1 + "/surveys/date/{date}/", controllers.GetSurveysByDate).Methods("GET")
	surveyRouter.HandleFunc(v1 + "/surveys/date/range/{from}/{to}/", controllers.GetSurveysByDateRange).Methods("GET")
	surveyRouter.HandleFunc(v1 + "/surveys/{id}", controllers.GetSurveyByID).Methods("GET")
	surveyRouter.HandleFunc(v1 + "/surveys/client/{name}/", controllers.GetSurveyByClient).Methods("GET")
	router.PathPrefix(v1 + "/surveys").Handler(gokyl.New(
		//gokyl.HandlerFunc(common.Authorize),
		gokyl.Wrap(surveyRouter),
	))

	return router
}
