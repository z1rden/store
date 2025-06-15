package kafka_admin

import "github.com/IBM/sarama"

func NewTopicDetail(opts ...TopicOption) *sarama.TopicDetail {
	topicDetail := &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
		ConfigEntries:     map[string]*string{},
	}

	for _, opt := range opts {
		opt(topicDetail)
	}

	return topicDetail
}
