package defs

// request 客户端发给服务端的
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

type NewComment struct {
	AuthorId int	`json:"author_id"`
	Content string `json:"content"`
}

type NewVideo struct {
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
}

// response 服务端发给客户端的
type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

type SignedIn struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

type UserSession struct {
	UserName string `json:"user_name"`
	SessionId string `json:"session_id"`
}

type UserInfo struct {
	Id int `json:"id"`
	Username string	`json:"user_name"`
}

type Comments struct {
	Comments []*Comment	`json:"comments"`
}

type VideosInfo struct {
	Videos []*VideoInfo	`json:"videos"`
}

// data model
type User struct {
	Id int
	LoginName string
	Pwd string
}

type VideoInfo struct {
	Id string	`json:"id"`
	AuthorId int 	`json:"author_id"`
	Name string `json:"name"`
	DisplayCTime string `json:"display_ctime"`
}

type Comment struct {
	Id string `json:"id"`
	VideoId string `json:"video_id"`
	Author string `json:"author"`
	Content string `json:"content"`
}

type SimpleSession struct {
	Username string	// 登录名
	TTL int64
}
