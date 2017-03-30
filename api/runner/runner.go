package runner

import (
	"github.com/moehandi/gokyl"
	"github.com/moehandi/campaign/api/routers"
	"github.com/moehandi/campaign/api/common"
	"net/http"
	"github.com/Sirupsen/logrus"
	"runtime"
	"log"
)

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)
	// Use all CPU Cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func Start() {
	router := routers.InitRoutes()

	h := gokyl.Classic()
	h.UseHandler(router)

	server := &http.Server{
		// // ":"+common.AppConfig.Port
		Addr:    common.AppConfig.Server,
		Handler: h,
	}

	logrus.Info("api runner on: ", server.Addr)
	server.ListenAndServe()
}
