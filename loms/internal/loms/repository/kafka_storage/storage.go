package kafka_storage

import (
	"context"
	"loms/internal/loms/db"
	"loms/internal/loms/repository/kafka_storage/sqlc"
)

type Storage interface {
	SendMessages(ctx context.Context, callback func(ctx context.Context, message *sqlc.KafkaOutbox) error) error
}

type storage struct {
	ctx      context.Context
	dbClient db.Client
}

func NewStorage(ctx context.Context, dbClient db.Client) Storage {
	return &storage{
		ctx:      ctx,
		dbClient: dbClient,
	}
}
