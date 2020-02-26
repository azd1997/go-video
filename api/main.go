package main

import (
	"github.com/azd1997/go-video/api/session"
	"net/http"

	"github.com/julienschmidt/httprouter"
)



func main() {
	prepare()	// 启动时需要现将数据库已有的session加载到内存中
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	// 用户相关
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	router.GET("/user/:username", GetUserInfo)

	// 视频相关
	router.POST("/user/:username/videos", AddNewVideo)
	router.GET("/user/:username/videos", ListAllVideos)
	router.DELETE("/user/:username/videos/:vid-id", DeleteVideo)

	// 评论相关
	router.POST("/videos/:vid-id/comments", PostComment)
	router.GET("/videos/:vid-id/comments", ShowComments)

	return router
}

func prepare() {
	session.LoadSessionFromDB()
}


// main -> middlerware -> defs(message,err) -> handlers -> dbops -> response
