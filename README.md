## Getting started
Run follow commands into `chat` root directory

Some services use Rabbitmq or Mongodb as dependency, make sure those are running before to run the chat services.

### Run dependencies
Mongodb and Rabbitmq runs in background with docker

`make run-dependencies`

### Run users service
Open new terminal and run user service

`make run-users-service`

### Run chat websocket server

`make run-websocket-server`

### Run bot
Open new terminal and run chat bot

`make run-chat-bot`

### Run front end
Front end was made with React, truly I'm not a front end developer and very far away of that, it's really basic.

`cd frontend && npm install`

then into frontend directory

`npm start dev`
