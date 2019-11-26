package main

import (
	"github.com/azd1997/go-video/scheduler/dbops"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	vid := p.ByName("vid-id")
	if len(vid) == 0 {
		sendResponse(w, http.StatusBadRequest, "video id should not be empty")
		return
	}

	if err := dbops.AddVideoDeleteRecord(vid); err != nil {
		sendResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	sendResponse(w, http.StatusOK, "delete task commit successfully")
	return
}
