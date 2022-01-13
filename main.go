package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker(workerId int, msg chan int) {
	for m := range msg {
		fmt.Printf("WorkerID: %v Msg: %v\n", workerId, m)
		time.Sleep(time.Second)
	}
}

func main() {
	queue := make(chan int)

	for w := 0; w <= runtime.NumCPU(); w++ {
		go worker(w, queue)
	}

	for i := 0; i <= 10; i++ {
		queue <- i
	}
}
