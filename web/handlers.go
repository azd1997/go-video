package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type SignInOrUpPage struct {
	Name string
}

type UserHomePage struct {
	Name string
}

func signInOrUpHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// go http库会在接收客户端请求时将客户端cookie也包到请求中
	// 如果我们发现cookie中的username和session id这两个任意一个取不到（取不到就会有“no cookie”的错误）
	// 就给用户看最原始的页面，提示需要注册或者登录；
	// 如果都去到了且是合法的，那么就直接重定向到这个用户的主页去
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		pageVar := &SignInOrUpPage{Name:"xxx"}

		t, err := template.ParseFiles(TemplateDir + Sign_In_Up)
		if err != nil {
			log.Printf("Parse template: %s: %s\n", err, TemplateDir + Sign_In_Up)
			return
		}

		t.Execute(w, pageVar)
		return
	}

	if len(cname.Value) != 0 && len(sid.Value) != 0 {		// 这里没有做很详细的检查
		http.Redirect(w, r, "/home", http.StatusFound)	// 重定向交给userHomeHandler处理
		return
	}
}

func userHomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 工作逻辑和登录相反，如果没相应cookie则跳转至登录页面

	cname, err1 := r.Cookie("username")
	_, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// 进入用户主页一定是填过表单的，所以，要读这个表单得到一些信息

	// 这里直接读表单数据而不作检验，是因为这部分检验交给前段去调用后端api检查，这样做有利于服务的解耦
	fname := r.FormValue("username")

	var pageVar *UserHomePage
	if len(cname.Value) != 0 {	// 说明cookie中有 -> 说明登陆过了
		pageVar = &UserHomePage{Name:cname.Value}
	} else if len(fname) != 0 {	// 之前没登陆过，则尝试从提交的表单中数据去登录
		pageVar = &UserHomePage{Name:fname}
	}

	t, err := template.ParseFiles(TemplateDir + Home)
	if err != nil {
		log.Printf("Parse template: %s: %s\n", err, TemplateDir + Home)
		return
	}

	t.Execute(w, pageVar)
}

func apiHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	// 如果请求方法不是POST直接驳回
	if r.Method != http.MethodPost {
		resp, _ := json.Marshal(ErrRequestNotRecognized)
		io.WriteString(w, string(resp))
		return
	}

	// request body 反json
	request, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	apiBody := &ApiBody{}
	if err := json.Unmarshal(request, apiBody); err != nil {
		resp, _ := json.Marshal(ErrRequestBodyParseFailed)
		io.WriteString(w, string(resp))
		return
	}

	// 真正去处理api转发
	apiRequest(apiBody, w, r)
}


func proxyUploadHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	u, _ := url.Parse("http://127.0.0.1:9000/")	// TODO: 不建议硬编码到程序，而应该变成配置
	proxy := httputil.NewSingleHostReverseProxy(u)		// 实现了域名转换，其他全都不动
	proxy.ServeHTTP(w, r)
}
