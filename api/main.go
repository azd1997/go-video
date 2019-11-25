package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)



func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)

	return router
}

// main -> middlerware -> defs(message,err) -> handlers -> dbops -> response
