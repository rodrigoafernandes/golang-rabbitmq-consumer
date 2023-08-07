package main

import (
	"dev.azure.com/hiae-mda/SRE_MDA/SRE.MDA.GOLANG.RABBITMQ.CONSUMER/internal/config"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	config.Setup()
	rabbitConnectionString := fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		config.CFG.RabbitmqUser, config.CFG.RabbitmqPassword, config.CFG.RabbitmqHost,
		config.CFG.RabbitmqPort, config.CFG.RabbitmqVHost)
	conn, err := amqp.Dial(rabbitConnectionString)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		config.CFG.RabbitmqQueue, // queue
		"",                       // consumer
		true,                     // auto-ack
		false,                    // exclusive
		false,                    // no-local
		false,                    // no-wait
		nil,                      // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s", d.Body)
		}
	}()

	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
