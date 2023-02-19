package main

import (
	"context"
	"event_service/internal/config"
	"event_service/internal/event/repositories"
	"event_service/internal/event/transport/kafka"
	"event_service/internal/event/usecases"
	logger "event_service/pkg/log"
	"github.com/Shopify/sarama"
	"github.com/jmoiron/sqlx"
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

	db, err := sqlx.Open("postgres", appConfig.PostgresConnLink)
	if err != nil {
		return err
	}
	logrus.Info("init DB... OK")

	repo := repositories.NewPsqlRepository(db)
	useCase := usecases.NewUseCase(repo)

	saramaConfig := sarama.NewConfig()
	// Handle errors manually
	saramaConfig.Consumer.Return.Errors = true

	saramaClient, err := sarama.NewClient([]string{appConfig.KafkaAddr}, saramaConfig)
	if err != nil {
		return err
	}

	kh := kafka.New(saramaClient, useCase)

	ctx := context.TODO()

	return kh.Handle(ctx)
}
