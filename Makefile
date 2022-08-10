run-db:
	docker-compose -f docker/apps/mongodb/docker-compose.yaml up -d

run-rabbit-queue:
	docker-compose -f docker/apps/rabbitmq/docker-compose.yaml up -d

run-dependencies: run-db run-rabbit-queue

run-users-service:
	go run backend/src/users/cmd/main.go

run-websocket-server:
	go run backend/src/chat_websocket_server/cmd/main.go

run-chat-bot:
	go run backend/src/chat_bot_command/cmd/main.go