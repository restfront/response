package response

import "net/http"

func (resp *Response) Success(data any) {
	resp.writeResponse(http.StatusOK, data)
}

func (resp *Response) Ok(data any) {
	resp.writeResponse(http.StatusOK, data)
}

func (resp *Response) Created(data any) {
	resp.writeResponse(http.StatusCreated, data)
}

func (resp *Response) Accepted(data any) {
	resp.writeResponse(http.StatusAccepted, data)
}

func (resp *Response) NoContent() {
	resp.writeResponse(http.StatusNoContent, nil)
}
