package api

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_healthCheck(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		purpose      string
		route        string
		expectedCode int
	}{
		{
			purpose:      "should return 200",
			route:        "/api/health",
			expectedCode: 200,
		},
		{
			purpose:      "should return 400",
			route:        "/api/health2",
			expectedCode: 404,
		},
	}

	app := Setup()

	for _, tt := range tests {
		t.Run(tt.purpose, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.route, nil)

			res, _ := app.Test(req, -1)

			as.Equal(tt.expectedCode, res.StatusCode)
		})
	}
}

func TestUnit_handleFileUpload(t *testing.T) {
	as := assert.New(t)
	app := Setup()

	validRoute := "/api/upload"

	tests := []struct {
		purpose          string
		mock             func() *http.Request
		expectedCode     int
		expectedResponse string
	}{
		// {
		// 	purpose:          "should return 200",
		// 	route:            "/api/upload",
		// 	expectedCode:     200,
		// 	expectedResponse: "video uploaded successfully",
		// },
		{
			purpose: "should return 404",
			mock: func() *http.Request {
				return httptest.NewRequest(http.MethodPost, "/api/upload2", nil)
			},
			expectedCode:     404,
			expectedResponse: "Cannot POST /api/upload2",
		},
		{
			purpose: "should return 400: request has no multipart/form-data Content-Type",
			mock: func() *http.Request {
				req := httptest.NewRequest(http.MethodPost, validRoute, nil)
				return req
			},
			expectedCode:     400,
			expectedResponse: "error uploading video: request has no multipart/form-data Content-Type",
		},
		{
			purpose: "should return 400: failed to save file",
			mock: func() *http.Request {
				body := &bytes.Buffer{}
				writer := multipart.NewWriter(body)

				ioWriter, err := writer.CreateFormFile("file", "test.mp4")
				as.NoError(err)

				_, err = ioWriter.Write([]byte("hello world"))
				as.NoError(err)
				as.NoError(writer.Close())

				req := httptest.NewRequest(http.MethodPost, validRoute, body)
				req.Header.Set("Content-Type", writer.FormDataContentType())
				req.Header.Set("Content-Length", strconv.Itoa(len(body.Bytes())))
				return req
			},
			expectedCode: 400,
			// expectedResponse: "error saving video: failed to save file",
			expectedResponse: "error saving video",
		},
	}

	for _, tt := range tests {
		t.Run(tt.purpose, func(t *testing.T) {
			req := tt.mock()

			res, _ := app.Test(req, -1)

			bodyBytes, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			as.Equal(tt.expectedCode, res.StatusCode)
			as.Contains(string(bodyBytes), tt.expectedResponse)
		})
	}
}
