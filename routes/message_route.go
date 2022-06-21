package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"producer/models"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
)

const MESSAGE_POST_ROUTE = "/produce/msg"

func PostMessage(context *gin.Context) {
	log.Println("Invoked post user endpoint.")

	var message models.Message

	if err := context.BindJSON(&message); err != nil {
		log.Println(err)
	}
	log.Printf("Payload: %v", message)

	connLink := fmt.Sprintf("amqp://%v:%v@%v:%v/", rabbitUsername, rabbitPassword, rabbitHost, rabbitPort)

	//1. Create the connection to RMQ~
	conn, err := amqp091.Dial(connLink)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RMQ", err)
	}

	//Defer the close, when function finishes close the connection
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to create channel", err)
	}

	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"publisher", //queue name
		false,       // durable
		false,       //delete when unused
		false,       //exclusive
		false,       //no-wait ?
		nil,         //arguments
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to create queue", err)
	}

	msgPayload, err := json.Marshal(models.NewMessagePayload(message.Content, message.UID))

	if err != nil {
		log.Println("Failed to parse payload as JSON.")
	}

	err = channel.Publish(
		"",         //exchange
		queue.Name, //routing key
		false,      //mandatory
		false,      //immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        msgPayload,
		},
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to publish message", err)
	}

	log.Println("Produce succsessful.")
	context.String(http.StatusOK, "Produce successful.")
}
