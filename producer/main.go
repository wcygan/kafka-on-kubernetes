package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/wcygan/kafka-on-kubernetes/generated/go/packet/v1"
	"log"
	"time"
)

func main() {
	write()
}

func write() {
	w := &kafka.Writer{
		Addr:     kafka.TCP("example-cluster-kafka-bootstrap:9092"),
		Topic:    "packet",
		Balancer: &kafka.LeastBytes{},
	}

	// produce a packet every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// initialize packet number
	packetNumber := int64(1)

	for {
		select {
		case <-ticker.C:
			packet := packetv1.Packet{Number: packetNumber}
			log.Println("Producing packet: ", packet.Number)

			err := w.WriteMessages(context.Background(),
				kafka.Message{
					Key:   []byte("Key"),
					Value: []byte(packet.String()),
				},
			)

			if err != nil {
				log.Fatal("failed to write messages:", err)
			}

			// increment packet number
			packetNumber++
		}
	}
}
