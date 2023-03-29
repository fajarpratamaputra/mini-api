package consumer

import "github.com/ThreeDotsLabs/watermill/message"

type Service interface {
	InsertData(payload message.Payload) error
	DeleteData(payload message.Payload) error
}

type Repository interface {
	LikeMongo(payload message.Payload) error
	UnLikeMongo(payload message.Payload) error
}
