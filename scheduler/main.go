package main

import (
	"github.com/azd1997/go-video/scheduler/taskrunner"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	go taskrunner.Start()

	r := RegisterHandlers()
	http.ListenAndServe(":7000", r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)

	return router
}
