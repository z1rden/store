package kafka_storage

import (
	"context"
	
	"fmt"
	"github.com/jackc/pgx/v5"

	"loms/internal/loms/repository/kafka_storage/sqlc"
)

func (s *storage) SendMessages(ctx context.Context, callback func(ctx context.Context, message *sqlc.KafkaOutbox) error) error {
	pool := s.dbClient.GetWriterPool()
	err := pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		queries := sqlc.New(pool).WithTx(tx)

		var err error
		messages, err := queries.GetMessages(ctx, sqlc.NullMessageStatusType{
			MessageStatusType: sqlc.MessageStatusTypeNew,
			Valid:             true,
		})
		if err != nil {
			return fmt.Errorf("failed to select outbox messages: %w", err)
		}

		for _, message := range messages {
			err := callback(ctx, &message)

			if err == nil {
				err = UpdateMessageStatus(ctx, tx, message.MessageID, sqlc.MessageStatusTypeSent)
			} else {
				err = UpdateMessageStatus(ctx, tx, message.MessageID, sqlc.MessageStatusTypeFailed)
			}

			if err != nil {
				return fmt.Errorf("failed to update outbox message status: %w", err)
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to start send messages: %w", err)
	}

	return nil
}
