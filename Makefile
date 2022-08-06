run-db:
	docker-compose -f docker/apps/mongodb/docker-compose.yaml up -d

build-users-app:
	docker build -t aflores04/chat/users -f docker/apps/users/Dockerfile .

run-users-app:
	docker run -d -p 3000:3000 aflores04/chat/users

run-app: run-db build-users-app run-users-app