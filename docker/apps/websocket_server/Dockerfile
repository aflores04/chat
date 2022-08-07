FROM golang:1.17-alpine

WORKDIR /app

COPY . /app

RUN apk update && apk add --no-progress --no-cache \
    bash build-base \
    ca-certificates \
    git \
    gcc \
    g++ \
    pkgconf \
    musl-dev

RUN cd /app && go mod download

RUN cd /app/src/websocket_server/cmd && GOOS=linux GOARCH=amd64 go build -a -v -tags musl main.go

EXPOSE 8010

CMD ["/app/src/websocket_server/cmd/main"]