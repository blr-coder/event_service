package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"time"
)

type Handler struct {
	kafkaClient sarama.Client
}

func New(client sarama.Client) *Handler {
	return &Handler{
		kafkaClient: client,
	}
}

func (h *Handler) Handle() error {

	consumer, err := sarama.NewConsumerFromClient(h.kafkaClient)
	if err != nil {
		return err
	}

	//TODO:  error handling
	defer consumer.Close()

	//TODO:  CONFIG!!!!!!
	topic := "quickstart"
	partition := 0

	partitionConsumer, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetOldest)
	if err != nil {
		return err
	}

	i := 0
	//var testEvent *usecase_models.CreateEventInput
	var testEvent TestEventStructFromProducer
	for ; ; i++ {
		msg := <-partitionConsumer.Messages()

		err = json.Unmarshal(msg.Value, &testEvent)
		if err != nil {
			panic(err)
		}
		spew.Dump(testEvent)

		if string(msg.Key) == "THE END" {
			break
		}
	}
	logrus.Infof("Finished. Received %d messages.\n", i)

	return nil
}

type TestEventStructFromProducer struct {
	Name     string    `json:"name"`
	Currency int64     `json:"currency"`
	SomeTime time.Time `json:"some_time"`
}
