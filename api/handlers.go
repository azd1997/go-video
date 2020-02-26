package main

import (
	"encoding/json"
	"github.com/azd1997/go-video/api/dbops"
	"github.com/azd1997/go-video/api/defs"
	"github.com/azd1997/go-video/api/session"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrDBError)
		return
	}

	// 写入正确，则创建session
	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{
		Success:   true,
		SessionId: id,
	}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), http.StatusCreated)		// 201
	}

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//username := p.ByName("user_name")
	//io.WriteString(w, fmt.Sprintf("%s\n", username))

	res, _ := ioutil.ReadAll(r.Body)
	//log.Printf("User [%s] requests to login in\n", res)
	log.Printf("%s\n", res)

	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		log.Printf("Login: %s\n", err)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}		// ubody存的是用户提交表单得来的名字密码信息，接下来要和数据库中的信息对比

	// 验证请求体
	uname := p.ByName("username")
	log.Printf("Login url name: %s\n", uname)
	log.Printf("Login body name: %s\n", ubody.Username)
	if uname != ubody.Username {
		sendErrorResponse(w, defs.ErrNotAuthUser)
		return
	}

	log.Printf("Login name: %s\n", ubody.Username)

	pwd, err := dbops.GetUserCredential(ubody.Username)
	log.Printf("Login pwd: %s\n", pwd)
	log.Printf("Login body pwd: %s\n", ubody.Pwd)
	if err != nil || len(pwd) == 0 || pwd != ubody.Pwd {
		sendErrorResponse(w, defs.ErrNotAuthUser)
		return
	}

	// 密码也匹配后要创建session并保存
	id := session.GenerateNewSessionId(ubody.Username)
	si := defs.SignedUp{
		Success:   true,
		SessionId: id,
	}
	if resp, err := json.Marshal(si); err != nil {
		sendErrorResponse(w, defs.ErrInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), http.StatusOK)
	}
}

func GetUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		log.Printf("GetUserInfo: Unauthorized user\n")
		return
	}

	uname := p.ByName("username")
	userinfo, err := dbops.GetUserInfo(uname)
	if err != nil {
		log.Printf("GetUserInfo: %v\n", err)
		sendErrorResponse(w, defs.ErrDBError)
		return
	}

	if resp, err := json.Marshal(userinfo); err != nil {
		sendErrorResponse(w, defs.ErrInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), http.StatusOK)
	}
}

func AddNewVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	if !ValidateUser(w, r) {
		log.Printf("GetUserInfo: Unauthorized user\n")
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	nvBody := &defs.NewVideo{}
	if err := json.Unmarshal(res, nvBody); err != nil {
		log.Printf("AddNewVideo: %v\n", err)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}

	video, err := dbops.AddNewVideo(nvBody.AuthorId, nvBody.Name)
	if err != nil {
		log.Printf("AddNewVideo: %v\n", err)
		sendErrorResponse(w, defs.ErrDBError)
		return
	}

	if resp, err := json.Marshal(video); err != nil {
		sendErrorResponse(w, defs.ErrInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), http.StatusCreated)
	}
}


func ListAllVideos(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	if !ValidateUser(w, r) {
		log.Printf("ListAllVideos: Unauthorized user\n")
		return
	}

	username := p.ByName("username")
	vs, err := dbops.ListVideoInfo(username, 0, int(time.Now().Unix()))
	if err != nil {
		log.Printf("ListAllVideos: %s\n", err)
		sendErrorResponse(w, defs.ErrDBError)
		return
	}

	vsi := defs.VideosInfo{Videos:vs}
	if resp, err := json.Marshal(vsi); err != nil {
		sendErrorResponse(w, defs.ErrInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), http.StatusCreated)
	}
}


func DeleteVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		log.Printf("DeleteVideo: Unauthorized user\n")
		return
	}

	vid := p.ByName("vid-id")
	if err := dbops.DeleteVideo(vid); err != nil {
		log.Printf("DeleteVideo: %s\n", err)
		sendErrorResponse(w, defs.ErrDBError)
		return
	}

	sendNormalResponse(w, "", http.StatusNoContent)
}

func PostComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		log.Printf("PostComment: Unauthorized user\n")
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	commentBody := defs.NewComment{}
	if err := json.Unmarshal(reqBody, commentBody); err != nil {
		log.Printf("PostComment: %s\n", err)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}

	vid := p.ByName("vid-id")
	if err := dbops.AddNewComment(vid, commentBody.AuthorId, commentBody.Content); err != nil {
		log.Printf("PostComment: %s\n", err)
		sendErrorResponse(w, defs.ErrDBError)
	} else {
		sendNormalResponse(w, "ok", http.StatusCreated)
	}
}

func ShowComments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		log.Printf("ShowComments: Unauthorized user\n")
		return
	}

	vid := p.ByName("vid-id")
	commentList, err := dbops.ListComments(vid, 0, int(time.Now().Unix()))
	if err != nil {
		log.Printf("ShowComments: %s\n", err)
		sendErrorResponse(w, defs.ErrDBError)
		return
	}

	cms := &defs.Comments{Comments:commentList}
	if resp, err := json.Marshal(cms); err != nil {
		log.Printf("ShowComments: %s\n", err)
		sendErrorResponse(w, defs.ErrInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), http.StatusOK)
	}
}