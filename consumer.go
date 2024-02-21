package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func StartKafka() {
	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "corinthians",
		GroupID:  "",
		MaxBytes: 1000,
	}

	reader := kafka.NewReader(conf)

	for {
		m, err := reader.FetchMessage((context.Background()))

		if err != nil {
			break
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		
		err = reader.CommitMessages(context.Background(), m)
		
		if err != nil {
			fmt.Println("Commiting message: ", string(m.Value))
			continue
		}

	}

	
	if err := reader.Close(); err != nil {
		fmt.Println("failed to close reader:", err)
	}

	
	
}
