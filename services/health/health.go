package health

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) HealthCheck() Response {
	return Response{
		Status: "OK",
	}
}
