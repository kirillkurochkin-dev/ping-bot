package main

import (
	"fmt"
	"os"
	"os/signal"
	"ping-bot/workerpool"
	"syscall"
	"time"
)

const (
	WorkersCount = 4
	Timeout      = time.Second * 6
	Interval     = time.Second * 10
)

var urls = []string{
	"https://telegram.org/",
	"https://ya.ru/",
	"https://google.com/",
	"https://golang.org/",
}

func main() {
	responses := make(chan workerpool.Response)
	pool := workerpool.NewPool(WorkersCount, Timeout, responses)

	pool.InitPool()

	go generateJobs(pool)
	go sendResponses(responses)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	pool.Stop()
}

func sendResponses(responses chan workerpool.Response) {
	go func() {
		for respons := range responses {
			fmt.Println(respons)
		}
	}()
}

func generateJobs(pool *workerpool.Pool) {
	for {
		for _, url := range urls {
			pool.AddJob(workerpool.Job{URL: url})
		}

		time.Sleep(Interval)
	}
}
