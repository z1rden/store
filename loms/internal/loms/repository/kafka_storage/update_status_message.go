package kafka_storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"loms/internal/loms/repository/kafka_storage/sqlc"
)

func UpdateMessageStatus(ctx context.Context, tx pgx.Tx, messageID pgtype.UUID, status sqlc.MessageStatusType) error {
	queries := sqlc.New(tx)

	err := queries.UpdateMessageStatus(ctx, sqlc.UpdateMessageStatusParams{
		Status: sqlc.NullMessageStatusType{
			MessageStatusType: status,
			Valid:             true,
		},
		MessageID: messageID,
	})

	if err != nil {
		return fmt.Errorf("failed to update message status: %v", err)
	}

	return nil
}
