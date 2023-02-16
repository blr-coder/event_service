package main

import (
	"event_service/internal/config"
	"event_service/internal/event/transport/kafka"
	logger "event_service/pkg/log"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	if err := runKafkaConsumer(); err != nil {
		log.Fatalln(err)
	}
}

func runKafkaConsumer() (err error) {

	logger.InitLogger()

	logrus.Info("RUN KAFKA WORKER...")

	appConfig, err := config.NewConfig("configs/dev_config.yaml")
	if err != nil {
		return err
	}
	logrus.Info("init config... OK")

	saramaConfig := sarama.NewConfig()
	// Handle errors manually
	saramaConfig.Consumer.Return.Errors = true

	saramaClient, err := sarama.NewClient([]string{appConfig.KafkaAddr}, saramaConfig)
	if err != nil {
		return err
	}

	kh := kafka.New(saramaClient)

	return kh.Handle()
}
