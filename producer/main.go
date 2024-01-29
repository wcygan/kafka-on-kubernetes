package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	packetv1 "github.com/wcygan/kafka-on-kubernetes/generated/go/packet/v1"
	"log"
	"time"
)

func main() {
	list()
}

func write() {
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "packet",
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func write2() {
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

func list() {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
}
