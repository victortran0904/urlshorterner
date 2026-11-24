package service

//go:generate mockery --name=HealthCheck --filename health.go --outpkg mocks

type HealthCheck interface {
	CheckHealth() (HealthResponse, error)
}

type healthCheck struct {
	serviceName string
	instanceID  string
}

type HealthResponse struct {
	Message     string `json:"message"`
	ServiceName string `json:"service_name"`
	InstanceID  string `json:"instance_id"`
}

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
