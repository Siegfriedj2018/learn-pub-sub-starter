package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

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

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Printf("\nReceived exit signal, shutting down server\n")
}
