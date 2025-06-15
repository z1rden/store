package kafka_storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"loms/internal/loms/repository/kafka_storage/sqlc"
)

func InsertMessage(ctx context.Context, tx pgx.Tx, message *Message) error {
	queries := sqlc.New(tx)

	err := queries.CreateMessage(ctx, sqlc.CreateMessageParams{
		Event: pgtype.Text{
			String: message.Event,
			Valid:  true,
		},
		EntityType: pgtype.Text{
			String: message.EntityType,
			Valid:  true,
		},
		EntityID: pgtype.Text{
			String: message.EntityID,
			Valid:  true,
		},
		Data: pgtype.Text{
			String: message.Data,
			Valid:  true,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to insert message in kafka outbox: %w", err)
	}

	return nil
}
