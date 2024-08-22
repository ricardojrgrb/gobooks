package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d got %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)

	go worker(1, ch)
	go worker(2, ch)

	for i := range 10 {
		ch <- i
	}
}
