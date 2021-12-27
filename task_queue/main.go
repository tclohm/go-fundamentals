package main

func main() () {
	q := taskq.NewQueue(3, 4, func(number interface{}) {
		log.Println(number)
	})
	go q.StartWorkers()

	for i := 0; i < 10; i++ {
		q.EnqueueJobBlocking(i)
	}

	time.Sleep(time.Second * 10)
}