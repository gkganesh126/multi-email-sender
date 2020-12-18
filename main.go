package main

import (
	"flag"
	"net/http"

	"github.com/gkganesh126/multi-email-sender/routers"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	flag.Set("logtostderr", "true")
	router := routers.InitRoutes()
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}
	glog.Info("Listening...")
	server.ListenAndServe()
}
