package main

import (
	"fmt"
	"log"
	"encoding/json"
	"github.com/streadway/amqp"
)


func main() {
	// RabbitMQ server connection information
	rabbitMQURL := "amqp://guest:guest@localhost:5672/"

	// Connect to RabbitMQ
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare the signup queue
	queueName := "signup_queue"
	_, err = ch.QueueDeclare(
		queueName, // Queue name
		false,     // Durable
		false,     // Delete when unused
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare the queue: %v", err)
	}

	// Create a SignupData object (replace with actual signup data)
	signup := SignupData{
		Username: "john_doe",
		Email:    "john@example.com",
	}

	// Convert SignupData to JSON
	body, err := json.Marshal(signup)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Publish the message to the signup queue
	err = ch.Publish(
		"",        // Exchange
		queueName, // Routing key (queue name)
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	fmt.Println("Signup message sent to the queue")
}
