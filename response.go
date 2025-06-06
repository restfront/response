package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	writer  http.ResponseWriter
	headers map[string]string
}

type ResponseResult struct {
	StatusCode int    `json:"code,omitempty"`
	Message    string `json:"message"`
} // @name ResponseResult

func New(w http.ResponseWriter) *Response {
	return &Response{
		writer: w,
		headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

func (resp *Response) AddHeader(key, value string) *Response {
	resp.writer.Header().Add(key, value)
	return resp
}

func (resp *Response) DeleteHeader(key string) *Response {
	resp.writer.Header().Del(key)
	return resp
}

func (resp *Response) writeResponse(statusCode int, data any) {
	resp.writeHeaders()
	resp.writeStatusCode(statusCode)

	if data != nil {
		switch v := data.(type) {
		case error:
			data = newResponseResult(statusCode, v.Error())
		case string:
			data = newResponseResult(statusCode, v)
		}

		_ = json.NewEncoder(resp.writer).Encode(data)
	}
}

func (resp *Response) writeHeaders() {
	for key, value := range resp.headers {
		resp.writer.Header().Set(key, value)
	}
}

func (resp *Response) writeStatusCode(statusCode int) {
	resp.writer.WriteHeader(statusCode)
}

func newResponseResult(statusCode int, message string) *ResponseResult {
	return &ResponseResult{
		StatusCode: statusCode,
		Message:    message,
	}
}
