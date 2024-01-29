package cmd

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	packetv1 "github.com/wcygan/kafka-on-kubernetes/generated/go/packet/v1"
	"log"
	"math/rand"
	"time"
)

var rngCmd = &cobra.Command{
	Use:   "rng",
	Short: "Write a packet with a random number to the 'packet' Kafka topic",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: this command doesn't work yet

		// Kafka broker address
		brokerAddress := "localhost:9092"

		// Create a new writer
		w := kafka.NewWriter(kafka.WriterConfig{
			Brokers:  []string{brokerAddress},
			Topic:    "packet",
			Balancer: &kafka.LeastBytes{},
		})

		// Always close the writer when done
		defer w.Close()

		// Generate a random number
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Int63()

		// Create a new packet
		packet := packetv1.Packet{Number: randomNumber}

		// Write the packet to the Kafka topic
		err := w.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte("Key"),
				Value: []byte(packet.String()),
			},
		)

		// Handle any errors
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

		// Log success
		fmt.Println("Successfully wrote packet: ", packet.Number)
	},
}
