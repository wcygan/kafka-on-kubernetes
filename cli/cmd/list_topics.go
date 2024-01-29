package cmd

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
)

var listTopicsCmd = &cobra.Command{
	Use:   "list-topics",
	Short: "List all topics in the Kafka cluster",
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}
