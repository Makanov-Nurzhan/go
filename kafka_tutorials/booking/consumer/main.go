package main

import (
	"booking/consumer/db"
	handler "booking/consumer/handlers"
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	postgres, err := db.NewPostgresDB("postgres", "5432", "user", "pass", "booking")
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer postgres.Close()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka:9092"},
		Topic:    "booking_created",
		GroupID:  "booking-consumer-group",
		MinBytes: 1,
		MaxBytes: 10e6,
		MaxWait:  1 * time.Second,
	})

	log.Println("Consumer started, waiting for messages...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("read error: %v", err)
			continue
		}

		if err := handler.ProcessBookingMessage(msg, postgres); err != nil {
			log.Printf("failed to process message: %v", err)
		}
	}
}
