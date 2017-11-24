package main

import (
	"fmt"
	"sync"
)

var jobChan chan int
var wg sync.WaitGroup

func worker(jobChan <-chan int) {
	defer wg.Done()
	for job := range jobChan {
		fmt.Printf("执行任务 %d \n", job)
	}
}

func main() {
	jobChan = make(chan int, 100)
	//入队
	for i := 1; i <= 10; i++ {
		jobChan <- i
	}

	wg.Add(1)
	close(jobChan)

	go worker(jobChan)
	wg.Wait()
}
