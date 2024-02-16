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
		m, err := reader.ReadMessage((context.Background()))

		if err != nil {
			fmt.Println("Some error ocurred", err)
			continue
		}

		fmt.Println("Message is: ", string(m.Value))
	}
}
