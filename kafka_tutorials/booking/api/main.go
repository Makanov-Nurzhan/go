package main

import (
	"booking/api/handlers"
	kafka2 "booking/api/kafka"
	"log"
	"net/http"
)

func main() {
	writer := kafka2.NewBookingProducer("kafka:9092", "booking_created")

	http.HandleFunc("/book", handlers.BookingHandler(writer))

	log.Println("API server is running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
