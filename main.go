package main

import (
	"fmt"
)



func main() {

	var messageAmount, batch, cmpId int
	fmt.Scanln(&messageAmount, &batch, &cmpId)

	//result := Generator(messageAmount, batch, cmpId )

	go StartKafka()
	
}