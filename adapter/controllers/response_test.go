package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRespondWithJSON(t *testing.T) {
	testCases := []struct {
		name       string
		statusCode int
		payload    interface{}
		wantStatus int
		wantHeader string
		wantBody   string
	}{
		{
			name:       "Valid JSON Payload",
			statusCode: http.StatusOK,
			payload:    map[string]string{"message": "success"},
			wantStatus: http.StatusOK,
			wantHeader: "application/json; charset=utf-8",
			wantBody:   `{"message":"success"}`,
		},
		{
			name:       "Integer Payload",
			statusCode: http.StatusCreated,
			payload:    123,
			wantStatus: http.StatusCreated,
			wantHeader: "application/json; charset=utf-8",
			wantBody:   `123`,
		},
		{
			name:       "String Payload",
			statusCode: http.StatusOK,
			payload:    "hello",
			wantStatus: http.StatusOK,
			wantHeader: "application/json; charset=utf-8",
			wantBody:   `"hello"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			ResponseJSON(rec, tc.statusCode, tc.payload)

			assert.Equal(t, tc.wantStatus, rec.Code)
			assert.Equal(t, tc.wantHeader, rec.Header().Get("Content-Type"))
			assert.JSONEq(t, tc.wantBody, rec.Body.String())
		})
	}
}

func TestRespondWithError(t *testing.T) {
	testCases := []struct {
		name       string
		statusCode int
		err        error
		wantStatus int
		wantHeader string
		wantBody   string
	}{
		{
			name:       "BadRequest Error",
			statusCode: http.StatusBadRequest,
			err:        errors.New("an error occurred"),
			wantStatus: http.StatusBadRequest,
			wantHeader: "application/json; charset=utf-8",
			wantBody:   `{"Code":"400","Message":"400 Bad Request: an error occurred"}`,
		},
		{
			name:       "InternalServerError",
			statusCode: http.StatusInternalServerError,
			err:        errors.New("internal server error"),
			wantStatus: http.StatusInternalServerError,
			wantHeader: "application/json; charset=utf-8",
			wantBody:   `{"Code":"500","Message":"500 Internal Server Error: internal server error"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			ResponseError(rec, tc.statusCode, tc.err)

			assert.Equal(t, tc.wantStatus, rec.Code)
			assert.Equal(t, tc.wantHeader, rec.Header().Get("Content-Type"))
			assert.JSONEq(t, tc.wantBody, rec.Body.String())
		})
	}
}
