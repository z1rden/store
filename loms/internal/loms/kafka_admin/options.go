package kafka_admin

import (
	"github.com/IBM/sarama"
	"strconv"
	"time"
)

type TopicOption func(topic *sarama.TopicDetail)

func WithNumPartitions(num int32) TopicOption {
	return func(topic *sarama.TopicDetail) {
		topic.NumPartitions = num
	}
}

func WithReplicationFactor(num int16) TopicOption {
	return func(topic *sarama.TopicDetail) {
		topic.ReplicationFactor = num
	}
}

func WithRetentionMSMinute(num int) TopicOption {
	return func(topic *sarama.TopicDetail) {
		ttl := time.Duration(num) * time.Minute
		retentionMs := ttl.Milliseconds()
		retentionMsStr := strconv.Itoa(int(retentionMs))

		topic.ConfigEntries["retention.ms"] = &retentionMsStr
	}
}
