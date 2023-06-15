package main

import (
	"event_service/internal/config"
	"event_service/internal/event/repositories"
	eventgrpc "event_service/internal/event/transport/grpc"
	"event_service/internal/event/usecases"
	"event_service/pkg/log"
	"fmt"
	"net"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func runEventApp() (err error) {
	log.InitLogger()

	logrus.Info("RUN APP...")

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

	eventServer := eventgrpc.NewEventServiceServer(useCase)
	eventTypeServer := eventgrpc.NewEventTypeServiceServer(useCase)
	reportServer := eventgrpc.NewReportServiceServer(useCase)

	grpcServer := eventgrpc.NewGRPCServer(eventServer, eventTypeServer, reportServer)
	logrus.Info("init grpcServer... OK")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", appConfig.AppPort))
	if err != nil {
		return err
	}
	logrus.Infof("listen grpc server in port: %s...", appConfig.AppPort)

	return grpcServer.Serve(listener)
}
