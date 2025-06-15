package kafka_admin

import (
	"github.com/IBM/sarama"
)

type Admin interface {
	CreateTopic(topicName string, opts ...TopicOption) error
	Close() error
}

type admin struct {
	clusterAdmin sarama.ClusterAdmin
}

func NewAdmin(addrs []string) (Admin, error) {
	config := sarama.NewConfig()

	clusterAdmin, err := sarama.NewClusterAdmin(addrs, config)
	if err != nil {
		return nil, err
	}

	return &admin{
		clusterAdmin: clusterAdmin,
	}, nil
}
