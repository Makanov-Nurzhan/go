// handler/booking.go
package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

type Booking struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Date   string `json:"date"`
	Guests int    `json:"guests"`
}

func ProcessBookingMessage(msg kafka.Message, db *sql.DB) error {
	var booking Booking
	if err := json.Unmarshal(msg.Value, &booking); err != nil {
		return err
	}

	_, err := db.ExecContext(context.Background(),
		`INSERT INTO bookings (id, name, phone, date, guests)
		 VALUES ($1, $2, $3, $4, $5)`,
		booking.ID, booking.Name, booking.Phone, booking.Date, booking.Guests)

	if err != nil {
		return err
	}

	log.Printf("Saved booking %s to DB\n", booking.ID)
	return nil
}
