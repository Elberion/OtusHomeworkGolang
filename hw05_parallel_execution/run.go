package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	tasksToRun := make(chan Task, len(tasks))
	notifyCh := make(chan error)
	defer close(notifyCh)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for _, task := range tasks {
			tasksToRun <- task
		}
		close(tasksToRun)
		wg.Done()
	}()
	wg.Add(1)
	go monitor(tasksToRun, len(tasks), m, n, notifyCh, wg)
	resError := <-notifyCh
	wg.Wait()
	return resError
}

func worker(tasks <-chan Task, results chan<- error, wg *sync.WaitGroup) {
	for task := range tasks {
		taskErr := task()
		results <- taskErr
	}
	wg.Done()
}

func monitor(tasks <-chan Task, tasksCount, errorLimit, workerCount int, notifyCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	readyTask := 0
	errorCount := 0
	taskToWorker := make(chan Task, workerCount)
	defer close(taskToWorker)
	results := make(chan error, tasksCount)

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		task, ok := <-tasks
		if ok {
			taskToWorker <- task
		}
		go worker(taskToWorker, results, wg)
	}

	for result := range results {
		if result != nil {
			errorCount++
			if errorCount >= errorLimit && errorLimit >= 0 {
				notifyCh <- ErrErrorsLimitExceeded
				return
			}
		}
		readyTask++
		if readyTask >= tasksCount {
			notifyCh <- nil
			return
		}

		task, ok := <-tasks
		if ok {
			taskToWorker <- task
		}
	}

}
