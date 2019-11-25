package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	return middleWareHandler{r:r}
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// 目的就是在真正接收HTTP服务之前先检查一些东西

	// 1. 检查session
	validateUserSession(r)

	// 2. 调用httprouter.Router的ServeHTTP方法。
	m.r.ServeHTTP(w, r)
}
