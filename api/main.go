package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)

	return router
}

// 测试当前服务器结果
//(base) eiger@eiger-ThinkPad-X1-Carbon-3rd:~/gopath-default/src/github.com/azd1997/go-video$ curl -X POST localhost:8000/user
//Create User Handler


