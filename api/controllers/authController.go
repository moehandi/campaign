package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/moehandi/campaign/api/common"
	"github.com/moehandi/campaign/api/data"
	"github.com/moehandi/campaign/api/models"
	"fmt"
)

// TODO AUthRegister Systematicaly
// Handler for HTTP Post - "/auth/register"
//func AuthRegister(w http.ResponseWriter, r *http.Request) {
//	var dataResource UserResource
//	// Decode the incoming User json
//	err := json.NewDecoder(r.Body).Decode(&dataResource)
//	if err != nil {
//		common.ShowAppError(
//			w,
//			err,
//			"Invalid User data",
//			500,
//		)
//		return
//	}
//	user := &dataResource.Data
//	context := NewContext()
//	defer context.Close()
//	col := context.DbCollection("users")
//	repo := &data.UserRepository{C: col}
//	// Insert User document
//	repo.CreateUser(user)
//	// Clean-up the hashpassword to eliminate it from response JSON
//	user.HashPassword = nil
//	j, err := json.Marshal(UserResource{Data: *user})
//	if err != nil {
//		common.ShowAppError(
//			w,
//			err,
//			"An unexpected error has occurred",
//			500,
//		)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusCreated)
//	w.Write(j)
//}

// TODO NEXT MUST BE GENERATE FOR NEW SIGN UP
func AuthRegister(w http.ResponseWriter, r *http.Request) {
	var authModel models.Auth
	// Decode the incoming AUth json
	err := json.NewDecoder(r.Body).Decode(&authModel)
	if err != nil {
		common.ShowAppError(
			w,
			err,
			"Invalid data",
			500,
		)
		return
	}
	auth := &authModel
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("auth")
	repo := &data.AuthRepository{C: col}
	// Insert User document
	err = repo.Register(auth)
	if err != nil {
		fmt.Println("DB Register auth Error: ", err)
		common.ShowAppError(
			w,
			err,
			"Invalid data",
			500,
		)
		return
	}

	common.ShowAppSuccess(w,  *auth, http.StatusOK)
}

// Handler for HTTP Post - "/auth/token"
func AuthToken(w http.ResponseWriter, r *http.Request) {
	var authModel TokenModel
	var token string
	// Decode the incoming Login json
	err := json.NewDecoder(r.Body).Decode(&authModel)
	if err != nil {
		common.ShowAppError(
			w,
			err,
			"Invalid Login data",
			500,
		)
		return
	}
	model := authModel
	authLogin := models.Auth{
		SenderId:    model.SenderId,
		SecureKey: model.SecureKey,
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("auth")
	repo := &data.AuthRepository{C: col}

	// Authenticate the login auth
	fmt.Println("Authenticate the Login auth:" + authLogin.SenderId, authLogin.SecureKey)
	auth, err := repo.Authenticate(authLogin)
	if err != nil {
		common.ShowAppError(
			w,
			err,
			"Invalid login credentials",
			401,
		)
		return
	}
	// Generate JWT token
	token, err = common.GenerateJWT(auth.SenderId, "member")
	if err != nil {
		common.ShowAppError(
			w,
			err,
			"Eror while generating the access token",
			500,
		)
		return
	}

	// For Security Reasons
	// Clean-up the secure_key & secure_key_hash to eliminate it from response JSON
	auth.SecureKeyHash = nil
	auth.SecureKey = ""
	authUser := AuthModel{
		Auth:  auth,
		Token: token,
	}
	common.ShowAppSuccess(w,  authUser, http.StatusOK)
}
