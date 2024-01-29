package cmd

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"time"
)

var topic string

var readAllCmd = &cobra.Command{
	Use:   "read-all",
	Short: "Read all messages a Kafka topic",
	Run: func(cmd *cobra.Command, args []string) {
		// Kafka broker address
		brokerAddress := "localhost:9092"

		// Create a new reader
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:   []string{brokerAddress},
			Topic:     topic,
			Partition: 0,
			MinBytes:  10e3, // 10KB
			MaxBytes:  10e6, // 10MB
		})

		// Always close the reader when done
		defer reader.Close()

		// Context to control the reader's lifetime
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		for {
			// Read a message
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				break // Exit if we encounter an error
			}

			// Process the message
			fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
		}
	},
}

func init() {
	readAllCmd.PersistentFlags().StringVarP(&topic, "topic", "t", "", "Kafka topic to read from")
}
