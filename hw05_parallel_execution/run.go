package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	taskToRun := make(chan Task)
	results := make(chan error)
	notifyCh := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(n)
	go func() {
		for i := 1; i <= n; i++ {
			go worker(taskToRun, results, wg, notifyCh)
		}
		wg.Wait()
		close(results)
	}()
	go func() {
		defer close(taskToRun)
		for _, task := range tasks {
			select {
			case <-notifyCh:
				return
			default:
				taskToRun <- task
			}
		}
	}()
	errorCount := 0
	for result := range results {
		if result != nil {
			errorCount++
			if errorCount >= m && m > 0 {
				close(notifyCh)
				for range results {
					errorCount++
				}
				wg.Wait()
				return ErrErrorsLimitExceeded
			}
		}
	}
	return nil
}

func worker(tasks <-chan Task, results chan<- error, wg *sync.WaitGroup, notifyCh <-chan struct{}) {
	defer wg.Done()
	for task := range tasks {
		select {
		case <-notifyCh:
			return
		default:
			results <- task()
		}
	}
}
