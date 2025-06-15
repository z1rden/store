package kafka_storage

import "database/sql"

const (
	MessageStatusNew    string = "new"
	MessageStatusSent   string = "sent"
	MessageStatusFailed string = "failed"
)

type Message struct {
	ID         string
	CreatedAt  sql.NullTime
	Event      string
	EntityType string
	EntityID   string
	Status     string
	Data       string
}
