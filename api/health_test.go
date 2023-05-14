package api

import (
	"net/http"
	"net/http/httptest"
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
