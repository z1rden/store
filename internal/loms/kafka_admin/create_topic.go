package kafka_admin

func (a *admin) CreateTopic(topicName string, opts ...TopicOption) error {
	topicDetail := NewTopicDetail(opts...)

	ok, err := a.checkTopic(topicName)
	if err != nil {
		return err
	} else if !ok {
		err := a.clusterAdmin.CreateTopic(topicName, topicDetail, false)
		if err != nil {

			return err
		}
	}

	return nil
}

func (a *admin) checkTopic(topicName string) (bool, error) {
	list, err := a.clusterAdmin.ListTopics()
	if err != nil {
		return false, err
	}

	if _, ok := list[topicName]; !ok {
		return false, nil
	}

	return true, nil
}

func (a *admin) Close() error {
	err := a.clusterAdmin.Close()
	if err != nil {
		return err
	}

	return nil
}
