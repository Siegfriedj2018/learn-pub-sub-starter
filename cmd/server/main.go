package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Siegfriedj2018/learn-pub-sub-starter/internal/pubsub"
	"github.com/Siegfriedj2018/learn-pub-sub-starter/internal/routing"
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
		log.Printf("encountered an error: %v", err)
		connection.Close()
		os.Exit(2)
	}

	jsonPaused, err := json.Marshal(routing.PlayingState{
		IsPaused: true,
	})
	if err != nil {
		log.Printf("encountered an error: %v", err)
		connection.Close()
		os.Exit(3)
	}
	err = pubsub.PublishJSON(conChan, routing.ExchangePerilDirect, routing.PauseKey, jsonPaused)
	if err != nil {
		log.Printf("encountered an err: %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Printf("\nReceived exit signal, shutting down server\n")
}
