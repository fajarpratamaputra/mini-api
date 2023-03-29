package consumer

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) InsertData(payload message.Payload) error {
	return s.repository.LikeMongo(payload)
}

func (s *service) DeleteData(payload message.Payload) error {
	return s.repository.UnLikeMongo(payload)
}
