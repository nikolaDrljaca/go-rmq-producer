FROM golang:1-alpine AS build

RUN apk add --no-cache git

WORKDIR /src

RUN go mod init producer
RUN go get github.com/gin-gonic/gin
RUN go get github.com/rabbitmq/amqp091-go

COPY . /src

RUN go build .

FROM alpine as runtime

COPY --from=build /src/producer /app/producer

CMD [ "/app/producer" ]