package data

import (
	"github.com/lunny/log"
	"github.com/moehandi/campaign/api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type SurveyRepository struct {
	C *mgo.Collection
}

func (r *SurveyRepository) Create(survey *models.Survey) error {
	obj_id := bson.NewObjectId()
	survey.Id = obj_id
	survey.CreatedAt = time.Now()
	survey.UpdatedAt = time.Now()
	err := r.C.Insert(&survey)
	return err
}

func (r *SurveyRepository) GetByDate(by_date time.Time) []models.Survey {

	start := time.Date(by_date.Year(), by_date.Month(), by_date.Day(), 0, 0, 0, 0, time.UTC) // date start ex: 27.07.2017, 00:00:00
	end := time.Date(by_date.Year(), by_date.Month(), by_date.Day(), 23, 59, 59, 0, time.UTC) // date end ex : 27.07.2017, 23:59:59

	var surveys []models.Survey

	err := r.C.Find(
		bson.M{ "created_at": bson.M{ "$gte": start, "$lt": end }}).All(&surveys)

	if err != nil {
		log.Println("ERROR BY DATE", err)
	}
	return surveys
}

func (r *SurveyRepository) GetByDateRange(dateFrom, dateTo time.Time) []models.Survey {
	log.Println("DARI ", dateFrom.Year(), dateFrom.Month(), dateFrom.Day())
	fromDate := time.Date(dateFrom.Year(), dateFrom.Month(), dateFrom.Day(), 0, 0, 0, 0, time.UTC)
	toDate := time.Date(dateTo.Year(), dateTo.Month(), dateTo.Day(), 23, 59, 59, 0, time.UTC)

	var surveys []models.Survey

	//result := models.Survey{}
	err := r.C.Find(
		bson.M{
			"created_at": bson.M{
				"$gte" : fromDate,
				"$lte": toDate,
			},
		}).All(&surveys)

	if err != nil {

	}
	return surveys
}

func (r *SurveyRepository) GetAll() []models.Survey {
	var surveys []models.Survey
	iter := r.C.Find(nil).Iter()
	result := models.Survey{}
	for iter.Next(&result) {
		surveys = append(surveys, result)
	}
	return surveys
}

func (r *SurveyRepository) GetById(id string) (survey models.Survey, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&survey)
	return
}

func (r *SurveyRepository) GetByClient(name string) (survey []models.Survey, err error) {
	//err = r.C.FindId(bson.ObjectIdHex(name)).One(&survey)
	err = r.C.Find(bson.M{"client": name}).All(&survey)
	if err != nil {
		panic(err)
	}
	return
}
