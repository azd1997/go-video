package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// 代理
// 将客户端请求转换成真正干活的后台api


// web app
// proxy
// api

// 浏览器跨域
// 比如域名对应的是 127.0.0.1:8000，如果在jquery中去跳转 127.0.0.1:9000就引发跨域了，很多浏览器会阻止，但是可以设置一个allow名单
// 但是这样做不太安全

// proxy转发模式就是比如说 用户访问的是 127.0.0.1:8000/upload，到了client层这里，将它转发给127.0.0.1:9000/upload
// 而api模式就是 用户在给127.0.0.1:8000/api发请求时，附带json信息，比如
// {
//		url:"",
//		method:"",
//		message:""
//	}
// 再在client.go这里用http.CLient重新构造请求发给真正的后台api

// api透传不适用于原声http请求，比如我们的upload服务，就没办法这么做，所以还需要proxy


// http.Client 的用法就是声明全局变量，调用其request方法
var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}


func apiRequest(b *ApiBody, w http.ResponseWriter, r *http.Request) {

	switch b.Method {
	case http.MethodGet, http.MethodPost, http.MethodDelete:
		if err := clientRequestAndResponse(b, w, r); err != nil {
			log.Printf("apiRequest: %v\n", err)
			return
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad api request")
	}
}

// 构造代理请求，发给后端api服务，并获取响应，再把这个响应传给客户端
func clientRequestAndResponse(b *ApiBody, w http.ResponseWriter, r *http.Request) error {
	req, _ := http.NewRequest(b.Method, b.Url, nil)
	req.Header = r.Header		// 透传
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	normalResponse(w, resp)
	return nil
}

// 将后台api服务返回给api代理的应答交给responsewriter包装一下，再传回给客户端
func normalResponse(w http.ResponseWriter, resp *http.Response)  {
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		re, _ := json.Marshal(ErrInternalFaults)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, string(re))
		return
	}

	w.WriteHeader(resp.StatusCode)
	io.WriteString(w, string(respBody))
}