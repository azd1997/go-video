package defs

import "net/http"

type Err struct {
	Error string 	`json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrRequestBodyParseFailed = ErrorResponse{
		HttpSC: http.StatusBadRequest,	// 400
		Error:  Err{
			Error:"Request body is not correctly parsed",
			ErrorCode:"001",
		},
	}

	ErrNotAuthUser = ErrorResponse{
		HttpSC: http.StatusUnauthorized,	// 401
		Error:  Err{
			Error:"User authentication failed",
			ErrorCode:"002",
		},
	}

	ErrDBError = ErrorResponse{
		HttpSC: http.StatusInternalServerError,		// 500
		Error:  Err{
			Error:"DB ops failed",
			ErrorCode:"003",
		},
	}

	ErrInternalFaults = ErrorResponse{
		HttpSC: http.StatusInternalServerError,		// 500
		Error:  Err{
			Error:"Internal service error",
			ErrorCode:"004",
		},
	}
)
