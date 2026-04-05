package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub/PublishJson"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	connectionStr := "amqp://guest:guest@localhost:5672/"
	fmt.Println("Starting Peril server...")

	connection, err := amqp.Dial(connectionStr)
	if err != nil {
		log.Fatalf("error connecting to localhost, %v\n", err)
	}
	defer connection.Close()
	fmt.Printf("Connection to localhost successful\n")
	conChan, err := connection.Channel()
	if err != nil {
		log.Printf("encountered an error in the connection Channel: %v", err)
		connection.Close()
		os.Exit(2)
	}

	err = PublishJSON(conChan)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Printf("\nReceived exit signal, shutting down server\n")
}
