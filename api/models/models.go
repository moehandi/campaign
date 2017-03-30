package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	Auth struct {
		Id            bson.ObjectId `bson:"_id,omitempty" json:"id"`
		SenderId      string        `json:"sender_id" bson:"sender_id"`
		SecureKey     string        `json:"secure_key,omitempty" bson:"secure_key"`
		SecureKeyHash []byte        `json:"secure_key_hash,omitempty" bson:"secure_key_hash"`
		LastAccess    time.Time     `json:"last_access" bson:"last_access"`
	}

	Client struct {
		Id   bson.ObjectId `bson:"_id" json:"id"`
		Name string        `json:"name"`
	}

	Survey struct {
		Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name      string        `json:"name"`
		Client    string        `json:"client"`
		Email     string        `json:"email"`
		Phone     string        `json:"phone"`
		CreatedAt time.Time     `json:"created_at,omitempty" bson:"created_at"`
		UpdatedAt time.Time     `json:"updated_at,omitempty" bson:"updated_at"`
		//CustomFields map[string]interface{}
		CustomFields map[string]string `json:"custom_fields" bson:"custom_fields"`
	}

	MetaStructs struct {
		FieldKey   string `json:"field_key"`
		FieldValue string `json:"field_value"`
	}
)
