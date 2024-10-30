package outbox

import (
	"encoding/json"
	"time"

	"github.com/IBM/sarama"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/domain"
)

type Converter struct{}

func (c Converter) DomainEventsToMessage(events []domain.DeviceEvent, topic string) []*sarama.ProducerMessage {
	messages := make([]*sarama.ProducerMessage, 0, len(events))
	for _, event := range events {
		messages = append(messages, c.DomainEventToMessage(event, topic))
	}
	return messages
}

func (Converter) DomainEventToMessage(event domain.DeviceEvent, topic string) *sarama.ProducerMessage {
	data, err := json.Marshal(event)
	if err != nil {
		panic(err)
	}
	return &sarama.ProducerMessage{
		Topic:     topic,
		Key:       nil,
		Value:     sarama.ByteEncoder(data),
		Timestamp: time.Now(),
	}
}
