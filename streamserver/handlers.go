package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	videoFile := VIDEO_DIR + vid

	video, err := os.Open(videoFile)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	defer video.Close()

	w.Header().Set("Content-Type", "video/mp4")		// 浏览器直到它是mp4就会用播放视频的方式解析
	http.ServeContent(w, r, "", time.Now(), video)
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}