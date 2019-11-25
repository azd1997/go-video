package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	return middleWareHandler{r:r, l:NewConnLimit(cc)}
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// 目的就是在真正接收HTTP服务之前先检查一些东西. "劫持"

	// 1. 尝试获取token以连接（限流器使用）
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests))		// 429
		return
	}
	defer m.l.ReleaseConn()		// 记得连接退出时将token还回去

	// 2. 调用httprouter.Router的ServeHTTP方法。
	m.r.ServeHTTP(w, r)
}
