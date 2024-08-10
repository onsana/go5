package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

// Orange represents an orange with a size
type Orange struct {
	Size float64 `json:"size"`
}

// Basket represents the storage for classified oranges
type Basket struct {
	Small  int
	Medium int
	Large  int
}

func main() {
	// Configuring Kafka consumer
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("oranges", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start Kafka partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	basket := Basket{}

	// Channel to trigger periodic statistics display
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	// Processing messages from Kafka
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var orange Orange
			err := json.Unmarshal(msg.Value, &orange)
			if err != nil {
				log.Printf("Failed to unmarshal orange: %v", err)
				continue
			}

			// Classify the orange by size
			if orange.Size < 7 {
				basket.Small++
			} else if orange.Size < 10 {
				basket.Medium++
			} else {
				basket.Large++
			}

		case <-ticker.C:
			// Display statistics every 10 seconds
			fmt.Printf("Oranges: small=%d, medium=%d, large=%d\n", basket.Small, basket.Medium, basket.Large)
		}
	}
}
