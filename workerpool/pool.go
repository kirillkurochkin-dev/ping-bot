package workerpool

import (
	"log"
	"sync"
	"time"
)

type Pool struct {
	worker       *Worker
	workersCount int

	jobs      chan Job
	responses chan Response

	wg       *sync.WaitGroup
	isClosed bool
}

func NewPool(workersCount int, timeout time.Duration, responses chan Response) *Pool {
	return &Pool{
		worker:       NewWorker(timeout),
		workersCount: workersCount,
		jobs:         make(chan Job),
		responses:    responses,
		wg:           new(sync.WaitGroup),
	}
}

func (p *Pool) Stop() {
	p.isClosed = true
	close(p.jobs)
	p.wg.Wait()
}

func (p *Pool) AddJob(j Job) {
	if p.isClosed {
		return
	}

	p.jobs <- j
	p.wg.Add(1)
}

func (p *Pool) InitPool() {
	for i := 0; i < p.workersCount; i++ {
		go p.InitWorker(i)
	}
}

func (p *Pool) InitWorker(id int) {
	for job := range p.jobs {
		p.responses <- p.worker.makeRequest(job)
		p.wg.Done()
	}

	log.Printf("[worker ID %d] finished proccesing", id)
}
