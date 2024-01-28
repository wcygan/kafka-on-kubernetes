package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	packetv1 "github.com/wcygan/kafka-on-kubernetes/generated/go/packet/v1"
)

func main() {
	// create a new kafka writer
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "packet",
		Balancer: &kafka.LeastBytes{},
	})

	defer w.Close()

	// produce a packet every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// initialize packet number
	packetNumber := int64(1)

	for {
		select {
		case <-ticker.C:
			packet := packetv1.Packet{Number: packetNumber}
			err := w.WriteMessages(context.Background(),
				kafka.Message{
					Key:   []byte("Key"),
					Value: []byte(packet.String()),
				},
			)

			if err != nil {
				log.Fatal("failed to write messages:", err)
			}

			log.Println("Producing packet: ", packet.Number)

			// increment packet number
			packetNumber++
		}
	}
}
