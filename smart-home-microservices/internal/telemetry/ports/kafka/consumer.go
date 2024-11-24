package kafka

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/IBM/sarama"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/telemetry/config"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/telemetry/domain"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/telemetry/usecase"
)

const (
	TopicDeviceCreated = "devices.created"
	TopicDeviceUpdated = "devices.updated"
	TopicDeviceDeleted = "devices.deleted"
)

func NewConsumer(cfg config.Config, app *usecase.App) *Consumer {
	clientConfig := sarama.NewConfig()
	clientConfig.ClientID = cfg.KafkaClientID

	consumer, err := sarama.NewConsumerGroup(cfg.KafkaDNSs, cfg.KafkaConsumerGroupID, clientConfig)
	if err != nil {
		panic(err)
	}

	return &Consumer{
		handler:       &handler{app: app},
		consumerGroup: consumer,
	}
}

type Consumer struct {
	handler       *handler
	consumerGroup sarama.ConsumerGroup
}

// Consume blocks until error
func (c *Consumer) Consume(ctx context.Context) error {
	return c.consumerGroup.Consume(ctx, []string{TopicDeviceCreated, TopicDeviceUpdated, TopicDeviceDeleted}, c.handler)
}

type handler struct {
	app *usecase.App
}

func (h *handler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		event := new(DeviceEvent)
		err := json.Unmarshal(msg.Value, event)
		if err != nil {
			return err
		}

		switch claim.Topic() {
		case TopicDeviceDeleted:
			h.handleDeviceDeleted(session, event)
		}

		session.MarkMessage(msg, "ok")
	}
	session.Commit()
	return nil
}

func (h *handler) handleDeviceDeleted(session sarama.ConsumerGroupSession, event *DeviceEvent) {
	err := h.app.DeleteStates.Handle(session.Context(), domain.ID(event.DeviceID))
	if err != nil {
		slog.Error(err.Error())
		return
	}
}

func (h *handler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h *handler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

type DeviceEvent struct {
	DeviceID string `json:"device_id"`
}
