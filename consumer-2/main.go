package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	packetv1 "github.com/wcygan/kafka-on-kubernetes/generated/go/packet/v1"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{
			"kafka-controller-0.kafka-controller-headless.default.svc.cluster.local:9092",
			"kafka-controller-1.kafka-controller-headless.default.svc.cluster.local:9092",
			"kafka-controller-2.kafka-controller-headless.default.svc.cluster.local:9092",
		},
		Topic: "packet",
	})

	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("failed to read message:", err)
		}

		var packet packetv1.Packet
		if err := proto.Unmarshal(m.Value, &packet); err != nil {
			log.Println("ERROR: failed to unmarshal message:", err)
		}

		if packet.GetNumber()%2 == 0 {
			log.Println("Found even number: ", packet.GetNumber())
		}
	}
}
