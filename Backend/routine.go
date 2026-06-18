package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for job := range jobs {
// 		fmt.Printf("worker %v hanlde jobs %v\n", id, job)
// 		results <- job * 2
// 	}
// }

// func main() {
// 	jobs := make(chan int, 10)
// 	results := make(chan int, 10)
// 	var wg sync.WaitGroup
// 	var workerCount int = runtime.NumCPU()
// 	for i := 0; i < workerCount; i++ {
// 		wg.Add(1)
// 		go worker(i, jobs, results, &wg)
// 	}
// 	for i := 0; i < 10; i++ {
// 		jobs <- i
// 	}
// 	close(jobs)
// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for i := 0; i < 10; i++ {
// 		msg := <-results
// 		fmt.Println(msg)
// 	}
// }
