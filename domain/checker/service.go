package checker

import "context"

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) CheckDBService() error {
	return s.repository.CheckDB()
}

func (s *service) CheckMongoDB(ctx context.Context) error {
	return s.repository.CheckMongoDB(ctx)
}
