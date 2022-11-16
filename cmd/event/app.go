package main

import (
	"event_service/internal/config"
	eventgrpc "event_service/internal/event/transport/grpc"
	"event_service/pkg/log"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
)

func runEventApp() (err error) {
	log.InitLogger()

	logrus.Info("RUN APP...")

	appConfig, err := config.NewConfig("configs/dev_config.yaml")
	if err != nil {
		return err
	}
	logrus.Info("init config... OK")

	eventServer := eventgrpc.NewEventServiceServer()
	eventTypeServer := eventgrpc.NewEventTypeServiceServer()

	grpcServer := eventgrpc.NewGRPCServer(eventServer, eventTypeServer)
	logrus.Info("init grpcServer... OK")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", appConfig.AppPort))
	if err != nil {
		return err
	}
	logrus.Infof("listen grpc server in port: %s...", appConfig.AppPort)

	return grpcServer.Serve(listener)
}
