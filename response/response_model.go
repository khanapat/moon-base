package response

type Response struct {
	Code  uint64      `json:"code" example:"200"`
	Title string      `json:"title" example:"Register key success."`
	Data  interface{} `json:"data,omitempty"`
}

type ErrResponse struct {
	Code        uint64      `json:"code" example:"400"`
	Title       string      `json:"title" example:"Cannot register public key."`
	Description string      `json:"description" example:"Please contact administrator for more information."`
	Error       interface{} `json:"error,omitempty"`
}

func NewResponse(code uint64, title string, data interface{}) *Response {
	return &Response{
		Code:  code,
		Title: title,
		Data:  data,
	}
}

func NewErrResponse(code uint64, title, desc string, err interface{}) *ErrResponse {
	return &ErrResponse{
		Code:        code,
		Title:       title,
		Description: desc,
		Error:       err,
	}
}
