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

const POST_USER_ROUTE = "/produce/user"

var rabbitHost = "localhost"
var rabbitPort = "5672"
var rabbitUsername = "guest"
var rabbitPassword = "guest"

func PostUser(context *gin.Context) {
	log.Println("Invoked post user endpoint.")

	var user models.User

	if err := context.BindJSON(&user); err != nil {
		log.Println(err)
	}
	log.Printf("Payload: %v", user)

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

	userPayload, err := json.Marshal(models.UserPayload{
		Name:  user.Name,
		Email: user.Email,
	})

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
			Body:        userPayload,
		},
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to publish message", err)
	}

	log.Println("Produce succsessful.")
	context.String(http.StatusOK, "Produce successful.")
}
