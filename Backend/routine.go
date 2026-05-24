package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("worker %v xu ly job %v \n", id, job)
		time.Sleep(time.Second)
		result <- job + 1
	}
}

func test() {
	jobs := make(chan int, 10)
	reusults := make(chan int, 10)
	var wg sync.WaitGroup
	for i := 1; i < 3; i++ {
		wg.Add(1)
		go worker(i, jobs, reusults, &wg)
	}

	for i := 1; i <= 9; i++ {
		jobs <- i
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(reusults)
	}()
	for r := range reusults {
		fmt.Println("Ket qua: ", r)
	}
}
