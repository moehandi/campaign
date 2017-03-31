package main

import (
	"net/http"

	"github.com/moehandi/gokyl"
	"github.com/moehandi/campaign/api/common"
	//_ "github.com/moehandi/campaign/api/common"
	"github.com/moehandi/campaign/api/routers"
	"runtime"
	"github.com/Sirupsen/logrus"
	"log"
)

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)
	// Use all CPU Cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	// Get the mux router object
	router := routers.InitRoutes()

	h := gokyl.New()
	h.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: h,
	}

	logrus.Info("api main running on: ", server.Addr)
	server.ListenAndServe()
}
