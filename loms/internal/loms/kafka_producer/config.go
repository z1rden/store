package kafka_producer

import (
	"github.com/IBM/sarama"
	"time"
)

func NewConfig(options ...ProducerOption) *sarama.Config {
	config := sarama.NewConfig()

	// Продюсер ожидает, когда в кафке сообщение запишется в master и во все isync реплики.
	config.Producer.RequiredAcks = sarama.WaitForAll
	// Кафка при false гарантирует запись exactly once сообщения в себе при возможных сбоях.
	config.Producer.Idempotent = false

	// Количество попыток продюсера отправить данные в kafka, если что-то пошло не так.
	config.Producer.Retry.Max = 100
	// Время между повторными отправками данных.
	config.Producer.Retry.Backoff = 100 * time.Millisecond

	// Выбор партиции под запись у кафки напрямую связана с ключом сообщения.
	// Если сообщения с одним и тем же ключом, то они попадают в одну партицию.
	config.Producer.Partitioner = sarama.NewHashPartitioner

	// Уровень сжатия.
	config.Producer.CompressionLevel = sarama.CompressionLevelDefault
	// Кодировшик компрессии.
	config.Producer.Compression = sarama.CompressionGZIP

	// Максимальное количество открытых соединений к кафке.
	// Значение 1 гарантирует упорядоченность значений в партициях.
	config.Net.MaxOpenRequests = 1

	// В случае успешной отправки сообщение будет отправляться в канал successes.
	config.Producer.Return.Successes = true
	// В случае успешной отправки сообщение будет отправляться в канал errors.
	config.Producer.Return.Errors = true

	for _, option := range options {
		option(config)
	}

	return config
}
