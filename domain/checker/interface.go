package checker

import "context"

type Service interface {
	CheckDBService() error
	CheckMongoDB(ctx context.Context) error
}

type Repository interface {
	Bbb() error
	CheckDB() error
	CheckMongoDB(ctx context.Context) error
}
