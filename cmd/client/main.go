package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/Siegfriedj2018/learn-pub-sub-starter/internal/gamelogic"
	"github.com/Siegfriedj2018/learn-pub-sub-starter/internal/pubsub"
	"github.com/Siegfriedj2018/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	connectionStr := "amqp://guest:guest@localhost:5672/"
	fmt.Println("Starting Peril client...")

	connection, err := amqp.Dial(connectionStr)
	if err != nil {
		log.Fatalf("error connecting to localhost: %v\n", err)
	}

	defer connection.Close()
	fmt.Println("Connection to localhost successful")

	username, err := gamelogic.ClientWelcome()
	if err != nil {
		log.Fatalf("could not get username: %v", err)
	}

	_, _, _ = pubsub.DeclareAndBind(
		connection,
		routing.ExchangePerilDirect,
		strings.Join([]string{routing.PauseKey, username}, "."),
		routing.PauseKey,
		pubsub.SimpleQueueTransient,
	)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Printf("\nReceived exit signal, shutting down server\n")
}
