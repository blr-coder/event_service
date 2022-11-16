.PHONY: build
.DEFAULT_GOAL := build

lint:
	golangci-lint run ./... --timeout 5m0s

fix_lint:
	golangci-lint run ./... --fix

build_proto:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./api/grpc/*/*.proto

db_up:
	sudo docker-compose up -d
db_stop:
	sudo docker-compose stop

migrate_up:
	migrate -path migrations -database "postgres://event_service_db_user:event_service_db_user_pass@localhost:5404/event_service_db?sslmode=disable" -verbose up
migrate_down:
	migrate -path migrations -database "postgres://event_service_db_user:event_service_db_user_pass@localhost:5404/event_service_db?sslmode=disable" -verbose down 1