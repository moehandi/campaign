package runner

import (
	"github.com/moehandi/gokyl"
	"github.com/moehandi/campaign/api/routers"
	"github.com/moehandi/campaign/api/common"
	"net/http"
	"github.com/Sirupsen/logrus"
)


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
