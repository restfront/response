package response

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type customError struct {
	code int
	msg  string
}

func (e customError) HTTPStatusCode() int {
	return e.code
}

func (e customError) Message() string {
	return e.msg
}

func TestSuccess(t *testing.T) {
	tests := []struct {
		name         string
		input        any
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Success with string",
			input:        "Success message",
			expectedCode: http.StatusOK,
			expectedBody: `{"code":200,"message":"Success message"}`,
		},
		{
			name:         "Success with struct",
			input:        map[string]string{"key": "value"},
			expectedCode: http.StatusOK,
			expectedBody: `{"key":"value"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			resp := New(w)
			resp.Success(tt.input)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		})
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		name         string
		input        any
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Error with string",
			input:        "Something went wrong",
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"code":500,"message":"Something went wrong"}`,
		},
		{
			name:         "Error with error",
			input:        errors.New("internal error"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"code":500,"message":"internal error"}`,
		},
		{
			name:         "Error with custom appError",
			input:        customError{code: http.StatusBadRequest, msg: "Invalid input"},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"code":400,"message":"Invalid input"}`,
		},
		{
			name:         "Error with nil",
			input:        nil,
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"code":500,"message":"Internal Server Error"}`,
		},
		{
			name:         "Error with number",
			input:        42,
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"code":500,"message":"Internal Server Error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			resp := New(w)
			resp.Error(tt.input)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		})
	}
}

func TestBadRequest(t *testing.T) {
	tests := []struct {
		name         string
		input        any
		expectedCode int
		expectedBody string
	}{
		{
			name:         "BadRequest with string",
			input:        "Invalid request",
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"code":400,"message":"Invalid request"}`,
		},
		{
			name:         "BadRequest with struct",
			input:        map[string]string{"error": "Invalid request"},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"Invalid request"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			resp := New(w)
			resp.BadRequest(tt.input)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		})
	}
}

func TestUnauthorized(t *testing.T) {
	w := httptest.NewRecorder()
	resp := New(w)
	resp.Unauthorized("Unauthorized")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.JSONEq(t, `{"code":401,"message":"Unauthorized"}`, w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func TestForbidden(t *testing.T) {
	w := httptest.NewRecorder()
	resp := New(w)
	resp.Forbidden("Forbidden")

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.JSONEq(t, `{"code":403,"message":"Forbidden"}`, w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func TestNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	resp := New(w)
	resp.NotFound("Not Found")

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.JSONEq(t, `{"code":404,"message":"Not Found"}`, w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func TestMethodNotAllowed(t *testing.T) {
	w := httptest.NewRecorder()
	resp := New(w)
	resp.MethodNotAllowed("Method Not Allowed")

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.JSONEq(t, `{"code":405,"message":"Method Not Allowed"}`, w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func TestUnprocessableEntity(t *testing.T) {
	w := httptest.NewRecorder()
	resp := New(w)
	resp.UnprocessableEntity("Unprocessable Entity")

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.JSONEq(t, `{"code":422,"message":"Unprocessable Entity"}`, w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func TestInternalServerError(t *testing.T) {
	w := httptest.NewRecorder()
	resp := New(w)
	resp.InternalServerError("Internal Server Error")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"code":500,"message":"Internal Server Error"}`, w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func TestServiceUnavailable(t *testing.T) {
	w := httptest.NewRecorder()
	resp := New(w)
	resp.ServiceUnavailable("Service Unavailable")

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	assert.JSONEq(t, `{"code":503,"message":"Service Unavailable"}`, w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func TestTooManyRequests(t *testing.T) {
	w := httptest.NewRecorder()
	resp := New(w)
	resp.TooManyRequests("Too Many Requests")

	assert.Equal(t, http.StatusTooManyRequests, w.Code)
	assert.JSONEq(t, `{"code":429,"message":"Too Many Requests"}`, w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func TestOk(t *testing.T) {
	w := httptest.NewRecorder()
	resp := New(w)
	resp.Ok("OK")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"code":200,"message":"OK"}`, w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func TestCreated(t *testing.T) {
	w := httptest.NewRecorder()
	resp := New(w)
	resp.Created("Created")

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"code":201,"message":"Created"}`, w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func TestAccepted(t *testing.T) {
	w := httptest.NewRecorder()
	resp := New(w)
	resp.Accepted("Accepted")

	assert.Equal(t, http.StatusAccepted, w.Code)
	assert.JSONEq(t, `{"code":202,"message":"Accepted"}`, w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func TestNoContent(t *testing.T) {
	w := httptest.NewRecorder()
	resp := New(w)
	resp.NoContent()

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func TestAddHeader(t *testing.T) {
	w := httptest.NewRecorder()
	resp := New(w)
	resp.AddHeader("X-Custom-Header", "value")

	assert.Equal(t, "value", w.Header().Get("X-Custom-Header"))
}

func TestDeleteHeader(t *testing.T) {
	w := httptest.NewRecorder()
	resp := New(w)
	resp.AddHeader("X-Custom-Header", "value")
	resp.DeleteHeader("X-Custom-Header")

	assert.Empty(t, w.Header().Get("X-Custom-Header"))
}

func TestWriteResponse(t *testing.T) {
	tests := []struct {
		name         string
		statusCode   int
		data         any
		expectedCode int
		expectedBody string
	}{
		{
			name:         "WriteResponse with string",
			statusCode:   http.StatusOK,
			data:         "Success",
			expectedCode: http.StatusOK,
			expectedBody: `{"code":200,"message":"Success"}`,
		},
		{
			name:         "WriteResponse with error",
			statusCode:   http.StatusInternalServerError,
			data:         errors.New("error"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"code":500,"message":"error"}`,
		},
		{
			name:         "WriteResponse with struct",
			statusCode:   http.StatusOK,
			data:         map[string]string{"key": "value"},
			expectedCode: http.StatusOK,
			expectedBody: `{"key":"value"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			resp := New(w)
			resp.writeResponse(tt.statusCode, tt.data)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		})
	}
}
