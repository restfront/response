package response

import (
	"net/http"
)

type appError interface {
	HTTPStatusCode() int
	Message() string
}

func (resp *Response) Error(data any) {
	var statusCode int
	var message string

	if appErr, ok := data.(appError); ok {
		statusCode = appErr.HTTPStatusCode()
		message = appErr.Message()
	} else {
		switch v := data.(type) {
		case string:
			statusCode = http.StatusInternalServerError
			message = v
		case error:
			statusCode = http.StatusInternalServerError
			message = v.Error()
		default:
			statusCode = http.StatusInternalServerError
			message = http.StatusText(statusCode)
		}
	}

	resp.writeResponse(statusCode, newResponseResult(statusCode, message))
}

func (resp *Response) BadRequest(data any) {
	resp.writeResponse(http.StatusBadRequest, data)
}

func (resp *Response) Unauthorized(data any) {
	resp.writeResponse(http.StatusUnauthorized, data)
}

func (resp *Response) Forbidden(data any) {
	resp.writeResponse(http.StatusForbidden, data)
}

func (resp *Response) NotFound(data any) {
	resp.writeResponse(http.StatusNotFound, data)
}

func (resp *Response) MethodNotAllowed(data any) {
	resp.writeResponse(http.StatusMethodNotAllowed, data)
}

func (resp *Response) UnprocessableEntity(data any) {
	resp.writeResponse(http.StatusUnprocessableEntity, data)
}

func (resp *Response) InternalServerError(data any) {
	resp.writeResponse(http.StatusInternalServerError, data)
}

func (resp *Response) ServiceUnavailable(data any) {
	resp.writeResponse(http.StatusServiceUnavailable, data)
}

func (resp *Response) TooManyRequests(data any) {
	resp.writeResponse(http.StatusTooManyRequests, data)
}
