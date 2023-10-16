package models

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseOK struct {
	Message string `json:"message"`
}

type FailureInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}

type DefaultResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
