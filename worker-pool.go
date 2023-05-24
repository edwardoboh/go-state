package main

import (
	"fmt"
	"time"
)

func worker(id int, receiveFrom <-chan int, sendTo chan<- int) {
	for i := range receiveFrom {
		fmt.Println("Starting task for worker: ", id, " | value: ", i)
		time.Sleep(time.Second)
		fmt.Println("Completed task for worker: ", id, " | value: ", i)
		sendTo <- i * 3
	}
}

func wp() {
	var chanSize int = 5
	sendChan := make(chan int, chanSize)
	recChan := make(chan int, chanSize)

	for i := 0; i < 3; i++ {
		go worker(i, recChan, sendChan)
	}

	for a := 0; a < chanSize; a++ {
		recChan <- a
	}
	close(recChan)

	for b := 0; b < chanSize; b++ {
		val := <-sendChan
		fmt.Println(val)
	}
}

/*
// NOTE - Rule of Thumb
1. Send-Only channel means you can only send a value to the channel: chan<-
2. Receive-Only channel means you can only receive a value from the channel: <-chan
*/
