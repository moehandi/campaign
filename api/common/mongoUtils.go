package common

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{AppConfig.MongoDBHost},
			Username: AppConfig.DBUser,
			Password: AppConfig.DBPwd,
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}
	return session
}
func createDbSession() {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{AppConfig.MongoDBHost},
		Username: AppConfig.DBUser,
		Password: AppConfig.DBPwd,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}
}

// Add indexes into MongoDB
func addIndexes() {
	var err error
	userIndex := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	authIndex := mgo.Index{
		Key:        []string{"sender_id"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	// Add indexes into MongoDB
	session := GetSession().Copy()
	defer session.Close()
	userCol := session.DB(AppConfig.MongoDBName).C("users")
	authCol := session.DB(AppConfig.MongoDBName).C("auth")

	err = userCol.EnsureIndex(userIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}

	err = authCol.EnsureIndex(authIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}

}
