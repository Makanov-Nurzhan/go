package handlers

import (
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"net/http"
)

type Booking struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Date   string `json:"date"` // ISO формат: 2025-05-05T19:00:00
	Guests int    `json:"guests"`
}

func BookingHandler(writer *kafka.Writer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var booking Booking
		if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}

		data, err := json.Marshal(booking)
		if err != nil {
			http.Error(w, "encoding error", http.StatusInternalServerError)
			return
		}

		err = writer.WriteMessages(r.Context(), kafka.Message{
			Value: data,
		})
		if err != nil {
			http.Error(w, "kafka error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Booking sent"))
	}
}
