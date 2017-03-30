package controllers

import (
	"github.com/moehandi/campaign/api/models"
	"time"
)

//Models for JSON resources
type (
	// TODO For Post - /auth/register
	RegisterPost struct {
		//Data models.Auth `json:"data"`
		SenderId  string `json:"sender_id"`
		SecureKey string `json:"secure_key"`
	}

	RegisterResponse struct {
		Data models.Auth `json:"data"`
	}

	// TODO For POST - /auth/token
	TokenResource struct {
		Data TokenModel `json:"data"`
	}

	// TODO Model for /auth/token
	TokenModel struct {
		SenderId      string `json:"sender_id"`
		SecureKey     string `json:"secure_key"`
		HashSecureKey string `json:"secure_key_hash"`
	}

	// TODO Response for Authorized POST /auth/token
	AuthResource struct {
		Data AuthModel `json:"data"`
	}

	// TODO Model for authorized with access token
	AuthModel struct {
		Auth  models.Auth `json:"auth"`
		Token string      `json:"token"`
	}

	SurveyResource struct {
		Data models.Survey `json:"data"`
	}

	SurveysResource struct {
		Data []models.Survey `json:"data"`
	}

	SurveyModel struct {
		Name         string            `json:"name"`
		Client       string            `json:"client"`
		Phone        string            `json:"phone"`
		Email        string            `json:"email"`
		CreatedOn    time.Time         `json:"created_on"`
		CustomFields map[string]string `json:"custom_fields" bson:"custom_fields"`
		//CustomFields map[string]interface{}
	}
)
