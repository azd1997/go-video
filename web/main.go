package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// 所谓的大前端就是包括 这里的web和template(传统意义的前端)

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":6000", r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", signInOrUpHandler)
	router.POST("/", signInOrUpHandler)

	router.GET("/home", userHomeHandler)
	router.POST("/home", userHomeHandler)

	router.POST("/api", apiHandler)		// api模式

	router.POST("upload/:vid-id", proxyUploadHandler)	// proxy模式

	router.ServeFiles("/statics/*filepath", http.Dir("./template"))	// 挂载文件夹

	return router
}