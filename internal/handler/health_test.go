package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/victortran0904/urlshortener/internal/service"
	health_svc_mocks "github.com/victortran0904/urlshortener/internal/service/mocks"
)

func Test_health_CheckHealth(t *testing.T) {
	testCases := []struct {
		name         string // description of this test case
		setupMockSvc func(ctx context.Context) *health_svc_mocks.HealthCheck

		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "normal case",
			setupMockSvc: func(ctx context.Context) *health_svc_mocks.HealthCheck {
				mockSvc := health_svc_mocks.NewHealthCheck(t)
				mockSvc.On("CheckHealth").Return(
					service.HealthResponse{
						Message:     "OK",
						ServiceName: "name",
						InstanceID:  "id",
					}, nil)
				return mockSvc
			},

			expectedStatus:   http.StatusOK,
			expectedResponse: `{"instance_id":"id","message":"OK","service_name":"name"}`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			rec := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(rec)

			mockSvc := tc.setupMockSvc(ginCtx)
			healthCheckHandler := NewHealthCheck(mockSvc)
			healthCheckHandler.CheckHealth(ginCtx)

			assert.JSONEq(t, tc.expectedResponse, rec.Body.String())
			assert.Equal(t, tc.expectedStatus, rec.Code)
		})
	}
}
