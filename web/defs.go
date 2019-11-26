package main

const (
	TemplateDir = "./templates/"

	Sign_In_Up = "sign_in_up.html"
	Home = "home.html"
)

// api透传格式定义。
type ApiBody struct {
	Url string `json:"url"`
	Method string `json:"method"`
	ReqBody string `json:"req_body"`
}

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

var (
	ErrRequestNotRecognized = Err{
		Error:     "api not recognized, bad request",
		ErrorCode: "001",
	}

	ErrRequestBodyParseFailed = Err{
		Error:     "request body is not correct",
		ErrorCode: "002",
	}

	ErrInternalFaults = Err{
		Error:     "internal service error",		// 不会暴露给客户，所以统一叫内部错误
		ErrorCode: "003",
	}
)
