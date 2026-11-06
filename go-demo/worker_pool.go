package main

import (
	"fmt"
	"sync"
)

type Job struct {
	ID int
}

func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("worker %d processing job %d\n", id, job.ID)
	}
}

func main() {
	const workerCount = 3
	const jobCount = 10

	jobs := make(chan Job)
	var wg sync.WaitGroup

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for i := 1; i <= jobCount; i++ {
		jobs <- Job{ID: i}
	}

	close(jobs)
	wg.Wait()

	fmt.Println("all jobs completed")
}
