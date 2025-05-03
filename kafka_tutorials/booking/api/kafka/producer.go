package kafka

import "github.com/segmentio/kafka-go"

func NewBookingProducer(brokerAddr, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(brokerAddr),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}
