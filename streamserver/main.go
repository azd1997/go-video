package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 2)		// 设为2方便测试时看出效果
	http.ListenAndServe(":9000", mh)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHandler)
	router.POST("/upload/:vid-id", uploadHandler)

	return router
}
