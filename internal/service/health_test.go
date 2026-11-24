package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		serviceName string
		instanceID  string
		expected    HealthResponse
	}{
		{
			name:        "normal case",
			serviceName: "urlshortner",
			instanceID:  "18cf5eaf-f401-45bf-a827-265b9c6c9333",
			expected: HealthResponse{
				Message:     "OK",
				ServiceName: "urlshortner",
				InstanceID:  "18cf5eaf-f401-45bf-a827-265b9c6c9333",
			},
		},
		{
			name:        "empty values",
			serviceName: "",
			instanceID:  "",
			expected: HealthResponse{
				Message:     "OK",
				ServiceName: "",
				InstanceID:  "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			svc := NewHealthCheck(tc.serviceName, tc.instanceID)

			got, err := svc.CheckHealth()

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, got)

		})
	}

}
