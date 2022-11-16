.PHONY: build
.DEFAULT_GOAL := build

lint:
	golangci-lint run ./... --timeout 5m0s

fix_lint:
	golangci-lint run ./... --fix

build_proto:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./api/grpc/*/*.proto
#db_up:
#	sudo docker-compose up -d
#db_stop:
#	sudo docker-compose stop
#run:
#	go run cmd/book/main.go --config /home/rabushko/CODE/GO_PROJECTS/book_grpc/configs/config.toml