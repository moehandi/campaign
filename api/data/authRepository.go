package data

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/moehandi/campaign/api/models"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"time"
)

type AuthRepository struct {
	C *mgo.Collection
}

func (r *AuthRepository) Register(auth *models.Auth) error {
	obj_id := bson.NewObjectId()
	auth.Id = obj_id

	hashkey, err := bcrypt.GenerateFromPassword([]byte(auth.SecureKey), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	auth.SecureKeyHash = hashkey
	// fmt.Printf("hashkey: %v\n", string(hashkey))
	// fmt.Printf("auth.HashSecureKey: %v \n", string(auth.SecureKeyHash))
	//clear the incoming text secure key
	//auth.SecureKey = ""
	auth.LastAccess = time.Now()
	//fmt.Println("AUTH","ID:" +auth.Id, "sender_id:" + auth.SenderId, "secure_key: " +auth.SecureKey, "secure_key_hash: "+string(auth.SecureKeyHash))
	err = r.C.Insert(&auth)
	return err
}

func (r *AuthRepository) Authenticate(auth models.Auth) (a models.Auth, err error) {

	err = r.C.Find(bson.M{"sender_id": auth.SenderId}).One(&a)
	if err != nil {
		return
	}

	// Validate secure key
	err = bcrypt.CompareHashAndPassword(a.SecureKeyHash, []byte(auth.SecureKey))

	if err != nil {

		errUpdate := updateLastAccess(auth,  r)
		if errUpdate != nil {
			fmt.Println("Error on Update last access")
			return
		}

		a = models.Auth{}
	}
	return
}
func updateLastAccess(auth models.Auth,  r *AuthRepository) error{
	now := time.Now()
	colQuerier := bson.M{"sender_id": auth.SenderId}
	change := bson.M{"$set": bson.M{"last_access": now}}
	err := r.C.Update(colQuerier, change)

	return err
}
