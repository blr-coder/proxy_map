package kafka

import (
	"errors"
	"github.com/Shopify/sarama"
)

type TopicDetails map[string]*sarama.TopicDetail

type Config struct {
	*sarama.Config
	Details TopicDetails
	Addr    []string
}

func (c *Config) setup(client sarama.Client) error {
	admin, err := sarama.NewClusterAdminFromClient(client)
	if err != nil {
		return err
	}

	for name, detail := range c.Details {
		err = admin.CreateTopic(name, detail, false)
		if errors.Is(err, sarama.ErrTopicAlreadyExists) {
			continue
		}
		if err != nil {
			return err
		}
	}

	return nil
}
