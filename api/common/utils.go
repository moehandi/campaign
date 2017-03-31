package common

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"net/http"
	"os"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
	successResponse struct {
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	}
	configuration struct {
		Server,
		//Port,
		MongoDBHost, DBUser, DBPwd, MongoDBName string
		LogLevel int
	}
)

// describe success on POST, PUT, DELETE
func ShowAppSuccess(w http.ResponseWriter, data interface{}, code int) {
	successObject := successResponse{
		Status: "success",
		Data:   data,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(successObject); err == nil {
		w.Write(j)
	}
}

func ShowAppError(w http.ResponseWriter, errorType error, message string, code int) {
	errObj := appError{
		Error:      errorType.Error(),
		Message:    message,
		HttpStatus: code,
	}
	Error.Printf("AppError]: %s\n", errorType)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}

// AppConfig holds the configuration values from config.json file
var AppConfig configuration

// Initialize AppConfig
func initConfig() {
	loadAppConfig()
}

// Reads config.json and decode into AppConfig
func loadAppConfig() {
	file, err := os.Open("config.json")
	defer file.Close()
	//fmt.Println("Initialize config with ./config.json")
	if err != nil {
		logrus.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		logrus.Fatalf("[loadAppConfig]: %s\n", err)
	}
}
