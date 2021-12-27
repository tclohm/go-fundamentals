package main

import (
	"log"
	"sync"
)

type JobCallBack func(job interface{})
type Queue struct {
	Workers int
	Capacity int
	JobQueue chan interface{}

	Wg *sync.WaitGroup
	QuitChan chan struct{}
	JobCallBack JobCallBack
}

func NewQueue(workers int, capacity int, jobCallBack JobCallBack) Queue {
	var wg sync.WaitGroup
	jobQueue := make(chan struct{})
	quit := make(chan struct{})
	return Queue{
		Workers: workers,
		JobQueue: jobQueue,
		JobCallBack: jobCallBack,
		Wg: &wg,
		QuitChan: quit,
	}
}

func (q *Queue) Stop() {
	q.QuitChan <- struct{}{}
}

func (q *Queue) EnqueueJobNonBlocking(job interface{}) bool {
	select {
	case q.JobQueue <- job:
		return true
	default:
		return false
	}
}

func (q *Queue) EnqueueJobBlocking(job interface{}) {
	q.JobQueue <- job
}

func (q *Queue) StartWorkers() {
	for i := 0; i < q.Workers; i++ {
		q.Wg.Add(1)
		go q.worker()
	}
	q.Wg.Wait()
}

func (q *Queue) worker() {
	defer q.Wg.Done()
	for {
		select {
		case <- q.QuitChan:
			log.Println("closing the workers")
			return
		case job := <- q.JobQueue:
			q.JobCallBack(job)
		}
	}
}