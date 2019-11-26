package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	videoFile := VIDEO_DIR + vid

	video, err := os.Open(videoFile)
	if err != nil {
		log.Printf("Open file error: %v\n", err)
		sendErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	defer video.Close()

	w.Header().Set("Content-Type", "video/mp4")		// 浏览器直到它是mp4就会用播放视频的方式解析
	http.ServeContent(w, r, "", time.Now(), video)
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}

	file, _, err := r.FormFile("file")		// 获取客户端提交的表单，选中“file”(就是上传视频的键)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v\n", err)
		sendErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	filename := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+filename, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v\n", err)
		sendErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	// 返回正确response
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")

	t.Execute(w, nil)
}