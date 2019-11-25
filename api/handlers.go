package main

import (
	"encoding/json"
	"fmt"
	"github.com/azd1997/go-video/api/dbops"
	"github.com/azd1997/go-video/api/defs"
	"github.com/azd1997/go-video/api/session"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
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
	username := p.ByName("user_name")
	io.WriteString(w, fmt.Sprintf("%s\n", username))
}
