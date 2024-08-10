package main

import (
	"encoding/json"
	"math/rand"
	"time"

	"log"
	"strconv"

	"github.com/IBM/sarama"
)

// Orange represents an orange with a random size
type Orange struct {
	Size float64 `json:"size"`
}

func main() {
	// Configuring Kafka producer
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka producer: %v", err)
	}
	defer producer.Close()

	rand.Seed(time.Now().UnixNano())

	// Sending oranges to Kafka in a loop
	for {
		orange := Orange{
			Size: rand.Float64()*10 + 5, // Random size between 5 and 15 cm
		}

		// Convert Orange struct to JSON
		orangeJSON, err := json.Marshal(orange)
		if err != nil {
			log.Printf("Failed to marshal orange: %v", err)
			continue
		}

		// Create a Kafka message
		msg := &sarama.ProducerMessage{
			Topic: "oranges",
			Key:   sarama.StringEncoder(strconv.Itoa(rand.Int())),
			Value: sarama.ByteEncoder(orangeJSON),
		}

		// Send the message
		_, _, err = producer.SendMessage(msg)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		} else {
			log.Printf("Sent orange with size: %.2f cm", orange.Size)
		}

		// Wait for a random time before sending the next message
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	}
}
