package main

import (
	"fmt"
)

var jobChan chan int

func worker(jobChan <-chan int) {
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

	close(jobChan)
	go worker(jobChan)

}
