package kafka

import (
	"context"
	"encoding/json"
	"event_service/internal/event/usecases"
	"event_service/internal/event/usecases/usecase_models"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"time"
)

type Handler struct {
	kafkaClient sarama.Client
	useCase     *usecases.UseCase
}

func New(client sarama.Client, useCase *usecases.UseCase) *Handler {
	return &Handler{
		kafkaClient: client,
		useCase:     useCase,
	}
}

func (h *Handler) Handle(ctx context.Context) error {

	consumer, err := sarama.NewConsumerFromClient(h.kafkaClient)
	if err != nil {
		return err
	}

	//TODO:  error handling
	defer consumer.Close()

	//TODO:  CONFIG!!!!!!
	topic := "quickstart"
	partition := 0

	partitionConsumer, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
	if err != nil {
		return err
	}

	i := 0
	var createEvent *usecase_models.CreateEventInput
	for ; ; i++ {
		msg := <-partitionConsumer.Messages()

		err = json.Unmarshal(msg.Value, &createEvent)
		//TODO:  error handling
		if err != nil {
			return err
		}

		event, err := h.useCase.Event.Create(ctx, createEvent)
		//TODO:  error handling
		if err != nil {
			return err
		}
		logrus.Infof("New event with type %s was created, %s", event.TypeTitle, event.CreatedAt.Format(time.RFC3339))

		if string(msg.Key) == "THE END" {
			i++
			break
		}
	}
	logrus.Infof("Finished. Received %d messages.\n", i)

	return nil
}
