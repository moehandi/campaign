package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lunny/log"
	"github.com/moehandi/campaign/api/common"
	"github.com/moehandi/campaign/api/data"
	"github.com/moehandi/campaign/api/models"
	"gopkg.in/mgo.v2"
	"net/http"
	"time"
)

func CreateSurvey(w http.ResponseWriter, r *http.Request) {

	var surveyModel models.Survey

	var surveyMap map[string]string

	err := json.NewDecoder(r.Body).Decode(&surveyMap)

	if err != nil {
		common.ShowAppError(w, err, "Invalid Survey data", 500)
		return
	}

	// Mandatory fields
	surveyModel.Name = surveyMap["name"]
	surveyModel.Email = surveyMap["email"]
	surveyModel.Client = surveyMap["client"]
	surveyModel.Phone = surveyMap["phone"]

	// check and initialize another custom fields
	for key, val := range surveyMap {
		if key != "email" && key != "name" && key != "phone" && key != "client" {
			if surveyModel.CustomFields == nil {
				surveyModel.CustomFields = make(map[string]string)
			}
			surveyModel.CustomFields[key] = val
		}
	}

	survey := &models.Survey{
		Name:         surveyModel.Name,
		Email:        surveyModel.Email,
		Phone:        surveyModel.Phone,
		Client:       surveyModel.Client,
		CreatedAt:    surveyModel.CreatedAt,
		CustomFields: surveyModel.CustomFields,
	}

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("surveys")

	// Insert a survey document
	repo := &data.SurveyRepository{C: col}
	repo.Create(survey)

	if err != nil {
		common.ShowAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	common.ShowAppSuccess(w, surveyModel, 201)
}

// GetTasks returns all Survey document
// Handler for HTTP Get - "/surveys"
func GetSurveys(w http.ResponseWriter, r *http.Request) {
	context := NewContext()

	defer context.Close()
	col := context.DbCollection("surveys")
	repo := &data.SurveyRepository{C: col}

	surveys := repo.GetAll()
	//j, err := json.Marshal(SurveysResource{Data: surveys})
	_, err := json.Marshal(SurveysResource{Data: surveys})
	if err != nil {
		common.ShowAppError(
			w, err, "An unexpected error has occurred", 500,
		)
		return
	}

	common.ShowAppSuccess(w, surveys, http.StatusOK)
	//w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(j)
}

func GetSurveysByDate(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	vars := mux.Vars(r)

	by_date := vars["date"]

	layout := "2006-01-02"
	//str := "2014-11-12T11:45:26.371Z"
	t, err := time.Parse(layout, by_date)

	if err != nil {
		fmt.Println(err)
	}

	defer context.Close()
	col := context.DbCollection("surveys")
	repo := &data.SurveyRepository{C: col}

	surveys := repo.GetByDate(t)
	//j, err := json.Marshal(SurveysResource{Data: surveys})
	//if err != nil {
	//	common.ShowAppError(
	//		w, err, "An unexpected error has occurred", 500,
	//	)
	//	return
	//}
	common.ShowAppSuccess(w, surveys, http.StatusOK)
	//w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(j)
}

func GetSurveysByDateRange(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	vars := mux.Vars(r)

	to := vars["to"]
	from := vars["from"]

	// see http://stackoverflow.com/questions/20530327/origin-of-mon-jan-2-150405-mst-2006-in-golang
	// see https://golang.org/src/time/format.go
	layout := "2006-01-02" // must be exact to this date layout
	//str := "2014-11-12T11:45:26.371Z"
	t, err := time.Parse(layout, to)
	f, err := time.Parse(layout, from)

	if err != nil {
		fmt.Println(err)
	}

	log.Println("FROM: ", from, " TO ", to)
	defer context.Close()
	col := context.DbCollection("surveys")
	repo := &data.SurveyRepository{C: col}

	surveys := repo.GetByDateRange(f, t)

	common.ShowAppSuccess(w, surveys, http.StatusOK)
}

func GetSurveyByID(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("surveys")
	repo := &data.SurveyRepository{C: col}
	note, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		common.ShowAppError(w, err, "An unexpected error has occurred", 500)
		return

	}
	common.ShowAppSuccess(w, note, http.StatusOK)
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//w.Write(j)
}

func GetSurveyByClient(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	client := vars["name"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("surveys")
	repo := &data.SurveyRepository{C: col}
	note, err := repo.GetByClient(client)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		common.ShowAppError(w, err, "An unexpected error has occurred", 500)
		return

	}

	common.ShowAppSuccess(w, note, http.StatusOK)
}
