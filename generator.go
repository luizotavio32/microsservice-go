package main

import (
	"strconv"
	"github.com/google/uuid"
)

type Message struct {
	CustomerId, CustomerName string
	Batch, CmpId int 
	
}

func Generator(messageAmount, batch, cmpId int) []Message {

	var messageSlice []Message 
	
	for i := 0; i < messageAmount; i++ {
		raw := Message{
			CustomerId: uuid.New().String(),
			CustomerName: "Customer_" + strconv.Itoa(i),
			Batch: batch,
			CmpId: cmpId,
		}

		messageSlice = append(messageSlice, raw)
	}

	return messageSlice
}

