package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID       int
	Duration time.Duration
}

type Result struct {
	Output string
}

// jobQueue <-chan Job -- канал, из которого читаем
// resultQueue chan<- Result -- канал, в который пишем
func worker(id int, jobQueue <-chan Job, resultQueue chan<- Result) {
	for job := range jobQueue {
		fmt.Printf("Worker: %d started job %d\n", id, job.ID)
		time.Sleep(job.Duration)
		resultQueue <- Result{Output: fmt.Sprintf("Worker: %d finished job %d\n", id, job.ID)}
		fmt.Printf("Worker: %d completed job %d\n", id, job.ID)
	}
}

func main() {
	const (
		numWorkers = 5
		numJob     = 10
	)
	jobQueue := make(chan Job, numJob)
	resultQueue := make(chan Result, numWorkers)

	//для синхронизации, чтобы дождаться, что все воркеры успешно доработали
	//нельзя копировать ни в коем случае
	//всегда передать по указателю или по ссылке
	wg := sync.WaitGroup{}

	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		//эта строка только для компилятора версии ниже 1.22, либо явно
		//передавать аргумент в функцию
		i := i

		//в горутины передается по ссылке
		go func() {
			defer wg.Done()
			worker(i, jobQueue, resultQueue)
		}()
	}

	for i := 1; i <= numJob; i++ {
		jobQueue <- Job{ID: i, Duration: time.Second * time.Duration(i)}

	}
	close(jobQueue)

	go func() {
		wg.Wait()
		close(resultQueue)
	}()

	for result := range resultQueue {
		fmt.Printf(result.Output)
	}

	fmt.Println("All job completed")

}
