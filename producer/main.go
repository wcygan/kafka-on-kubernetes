package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	packetv1 "github.com/wcygan/kafka-on-kubernetes/generated/go/packet/v1"
	"log"
	"math/rand"
	"net"
	"time"
)

func main() {
	w := &kafka.Writer{
		Addr:     kafka.TCP("example-cluster-kafka-bootstrap:9092"),
		Topic:    "packet",
		Balancer: &kafka.LeastBytes{},
		Transport: &kafka.Transport{
			Dial: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).DialContext,
		},
	}

	for {
		packet := packetv1.Packet{Number: rand.Int63()}
		log.Println("Producing packet: ", packet.Number)

		err := w.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte("Key"),
				Value: []byte(packet.String()),
			},
		)

		if err != nil {
			log.Fatal("failed to write messages:", err)
		} else {
			log.Println("Packet produced: ", packet.Number)
		}
	}
}
