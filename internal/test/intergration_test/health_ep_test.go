package intergration_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victortran0904/urlshortener/internal/api"
)

func Test_health_CheckHealth(t *testing.T) {
	testCases := []struct {
		name string // description of this test case

		setupTestRequest func() (*httptest.ResponseRecorder, *http.Request)

		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "normal case",

			setupTestRequest: func() (*httptest.ResponseRecorder, *http.Request) {
				req := httptest.NewRequest("GET", "/health", nil)
				rec := httptest.NewRecorder()
				return rec, req
			},

			expectedStatus:   http.StatusOK,
			expectedResponse: `{"message":"OK", "service_name":"","instance_id":""}`,
		},
		{
			name: "wrong endpoint",

			setupTestRequest: func() (*httptest.ResponseRecorder, *http.Request) {
				req := httptest.NewRequest("GET", "/health123", nil)
				rec := httptest.NewRecorder()
				return rec, req
			},

			expectedStatus:   http.StatusNotFound,
			expectedResponse: ``,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			testAPIEngine := api.NewEngine(&api.Config{})
			rec, req := tc.setupTestRequest()
			testAPIEngine.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			if tc.expectedResponse != "" {
				assert.JSONEq(t, tc.expectedResponse, rec.Body.String())
			}
		})
	}
}
