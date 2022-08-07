run-db:
	docker-compose -f docker/apps/mongodb/docker-compose.yaml up -d

run-rabbit-queue:
	docker-compose -f docker/apps/rabbitmq/docker-compose.yaml up -d

build-websocket-server:
	docker build -t aflores04/chat/websocket_server -f docker/apps/websocket_server/Dockerfile .

build-users-app:
	docker build -t aflores04/chat/users -f docker/apps/users/Dockerfile .

run-users-app:
	docker run -d -p 3001:3001 aflores04/chat/users

run-websocket-server:
	docker run -d -p 8010:8010 aflores04/chat/websocket_server

run-dependencies: run-db run-rabbit-queue

run-app: run-db run-rabbit-queue build-users-app build-websocket-server run-users-app run-websocket-server