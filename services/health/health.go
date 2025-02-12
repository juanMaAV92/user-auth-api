package health

type Service struct{}

type HealthService interface {
	HealthCheck() Response
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) HealthCheck() Response {
	return Response{
		Status: "OK",
	}
}
