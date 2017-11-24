package main

import (
	"fmt"
	"time"
)

var jobChan chan int

func worker(jobChan <-chan int, cancelChan <-chan struct{}) {
	for {
		select {
		case <-cancelChan:
			return
		case job := <-jobChan:
			fmt.Printf("执行任务 %d \n", job)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	jobChan = make(chan int, 100)
	//通过chan 取消操作
	cancelChan := make(chan struct{})
	//入队
	for i := 1; i <= 10; i++ {
		jobChan <- i
	}

	close(jobChan)

	go worker(jobChan, cancelChan)
	time.Sleep(2 * time.Second)
	//关闭chan
	close(cancelChan)
}
