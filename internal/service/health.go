// Package service provides health responses and random password generation.
package service

//go:generate mockery --name=HealthCheck --filename health.go --outpkg mocks

// HealthCheck reports this service's health status and identity.
type HealthCheck interface {
	// CheckHealth returns a health response or an error.
	CheckHealth() (HealthResponse, error)
}

type healthCheck struct {
	serviceName string
	instanceID  string
}

// HealthResponse is the JSON body returned by a successful health check.
type HealthResponse struct {
	Message     string `json:"message"`
	ServiceName string `json:"service_name"`
	InstanceID  string `json:"instance_id"`
}

// NewHealthCheck creates a health check with the supplied service name and instance ID.
// The check returns "OK", the supplied identity, and a nil error.
func NewHealthCheck(serviceName string, instanceID string) HealthCheck {
	return &healthCheck{
		serviceName: serviceName,
		instanceID:  instanceID,
	}
}

func (h *healthCheck) CheckHealth() (HealthResponse, error) {

	return HealthResponse{
		Message:     "OK",
		ServiceName: h.serviceName,
		InstanceID:  h.instanceID,
	}, nil

}
