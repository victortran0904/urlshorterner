package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	genpass_svc_mocks "github.com/victortran0904/urlshortener/internal/service/mocks"
)

func Test_genPass_GeneratePassword(t *testing.T) {
	testCases := []struct {
		name         string // description of this test case
		setupMockSvc func(ctx context.Context) *genpass_svc_mocks.GenPass

		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "normal case",
			setupMockSvc: func(ctx context.Context) *genpass_svc_mocks.GenPass {
				mockSvc := genpass_svc_mocks.NewGenPass(t)
				mockSvc.On("GeneratePassword", DefaultLen).Return("password", nil)
				return mockSvc
			},

			expectedStatus:   http.StatusOK,
			expectedResponse: `{"password":"password"}`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			rec := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(rec)

			mockSvc := tc.setupMockSvc(ginCtx)
			genPassHandler := NewGenPass(mockSvc)
			genPassHandler.GeneratePassword(ginCtx)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, tc.expectedResponse, rec.Body.String())
		})
	}
}
