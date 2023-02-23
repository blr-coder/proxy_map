package kafka

import (
	"github.com/Shopify/sarama"
	"time"
)

type Conn struct {
	client   sarama.Client
	producer sarama.SyncProducer
	consumer sarama.Consumer
}

func Open(config *Config) (*Conn, error) {
	config.Producer.Return.Successes = true
	saramaClient, err := sarama.NewClient(config.Addr, config.Config)
	if err != nil {
		return nil, err
	}

	if err = config.setup(saramaClient); err != nil {
		return nil, err
	}

	producer, err := sarama.NewSyncProducerFromClient(saramaClient)
	if err != nil {
		return nil, err
	}

	return &Conn{
		client:   saramaClient,
		producer: producer,
	}, nil
}

func (c *Conn) Close() error {
	return c.client.Close()
}

type Message struct {
	Topic string
	Key   string
	Bytes []byte
}

func (c *Conn) Send(message Message) error {
	msg := &sarama.ProducerMessage{
		Topic:     message.Topic,
		Key:       sarama.StringEncoder(message.Key),
		Value:     sarama.ByteEncoder(message.Bytes),
		Timestamp: time.Now(),
	}

	_, _, err := c.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}
