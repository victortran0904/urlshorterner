package service

type Health interface{
	CheckHealth() (map[string]string, error)
}

type health struct{
}

func NewHealth() Health{
	return &health{}
}

func (h *health) CheckHealth() (map[string]string, error) {
	return map[string]string{
		"message": "",
		"name": "",
		"id": "",
	}, nil
}

