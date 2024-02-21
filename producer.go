package main

import (
	"time"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func ProduceKafka() {

	// to produce messages
	topic := "corinthians"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		fmt.Print("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		fmt.Print("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		fmt.Print("failed to close writer:", err)
	}
	
}