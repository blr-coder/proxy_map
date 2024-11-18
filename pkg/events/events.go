package events

import (
	"context"
	"github.com/Shopify/sarama"
	"time"
)

type EventSender interface {
	Send(ctx context.Context, request *CreateEventRequest) error
}

type KafkaSender struct {
	kafkaProducer sarama.SyncProducer
}

func NewKafkaSender(producer sarama.SyncProducer) *KafkaSender {
	return &KafkaSender{
		kafkaProducer: producer,
	}
}

func (s *KafkaSender) Send(ctx context.Context, request *CreateEventRequest) error {

	msg := &sarama.ProducerMessage{
		Topic:     "quickstart",
		Key:       sarama.StringEncoder("new_event"),
		Value:     request,
		Headers:   nil,
		Metadata:  nil,
		Offset:    -1,
		Partition: 0,
		Timestamp: time.Now(),
	}

	// TODO: ???????
	_, _, err := s.kafkaProducer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}

type CreateEventRequest struct {
	TypeTitle   string `json:"type_title"`
	CampaignID  int64  `json:"campaign_id"`
	InsertionID int64  `json:"insertion_id"`
	UserID      int64  `json:"user_id"`
	Cost        *Cost  `json:"cost"`
}

func (c *CreateEventRequest) Encode() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CreateEventRequest) Length() int {
	//TODO implement me
	panic("implement me")
}

type Cost struct {
	Amount   uint64 `json:"amount"`
	Currency string `json:"currency"`
}
