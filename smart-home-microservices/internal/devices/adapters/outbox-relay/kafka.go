package outbox

import (
	"context"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/config"
	"github.com/ismashin/architecture-sprint-3/smart-home-microservices/internal/devices/domain"
)

const (
	TopicDeviceCreated = "devices.created"
	TopicDeviceUpdated = "devices.updated"
	TopicDeviceDeleted = "devices.deleted"
)

type KafkaPollingRelay struct {
	repo           domain.EventsOutbox
	messageBroker  sarama.SyncProducer
	pollingTimeout time.Duration
	triggerCh      chan struct{}
	converter      Converter
}

func NewKafkaPollingRelay(cfg config.Config, repo domain.EventsOutbox) *KafkaPollingRelay {
	producer, err := sarama.NewSyncProducer(cfg.KafkaDNSs, sarama.NewConfig())
	if err != nil {
		panic(err)
	}
	return &KafkaPollingRelay{
		repo:           repo,
		messageBroker:  producer,
		pollingTimeout: cfg.OutboxRelayPollingInterval,
		triggerCh:      make(chan struct{}, 100),
	}
}

func (r *KafkaPollingRelay) Start(ctx context.Context) {
	ticker := time.NewTicker(r.pollingTimeout)
	for {
		r.handle(ctx)

		select {
		case <-ticker.C:
			r.handle(ctx)
		case <-r.triggerCh:
			r.handle(ctx)
			ticker.Reset(r.pollingTimeout)
		case <-ctx.Done():
			return
		}
	}
}

func (r *KafkaPollingRelay) TriggerReadDeviceDeleted(ctx context.Context) {
	r.trigger()
}

func (r *KafkaPollingRelay) TriggerReadDeviceCreated(ctx context.Context) {
	r.trigger()
}

func (r *KafkaPollingRelay) TriggerReadDeviceUpdated(ctx context.Context) {
	r.trigger()
}

func (r *KafkaPollingRelay) trigger() {
	if len(r.triggerCh) < cap(r.triggerCh) {
		r.triggerCh <- struct{}{}
	}
}

type DeviceEvent struct {
	DeviceID string `json:"device_id"`
}

func (r *KafkaPollingRelay) handle(ctx context.Context) {
	wg := &sync.WaitGroup{}
	wg.Add(3)

	go r.handleCreate(ctx, wg)
	go r.handleUpdate(ctx, wg)
	go r.handleDelete(ctx, wg)
}

func (r *KafkaPollingRelay) handleCreate(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	events, err := r.repo.GetDeviceCreatedEvents(ctx)
	if err != nil {
		return
	}
	err = r.messageBroker.SendMessages(r.converter.DomainEventsToMessage(events, TopicDeviceCreated))
	if err != nil {
		return
	}
	err = r.repo.DeleteDeviceCreatedEvents(ctx)
	if err != nil {
		return
	}
}

func (r *KafkaPollingRelay) handleUpdate(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	events, err := r.repo.GetDeviceUpdatedEvents(ctx)
	if err != nil {
		return
	}
	err = r.messageBroker.SendMessages(r.converter.DomainEventsToMessage(events, TopicDeviceUpdated))
	if err != nil {
		return
	}
	err = r.repo.DeleteDeviceUpdatedEvents(ctx)
	if err != nil {
		return
	}
}

func (r *KafkaPollingRelay) handleDelete(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	events, err := r.repo.GetDeviceDeletedEvents(ctx)
	if err != nil {
		return
	}
	err = r.messageBroker.SendMessages(r.converter.DomainEventsToMessage(events, TopicDeviceDeleted))
	if err != nil {
		return
	}
	err = r.repo.DeleteDeviceDeletedEvents(ctx)
	if err != nil {
		return
	}
}
