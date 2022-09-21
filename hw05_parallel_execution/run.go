package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in workersCount goroutines and stops its work when receiving maxErrorsCount errors from tasks.
func Run(tasks []Task, workersCount, maxErrorsCount int) error {

	var errorsCount int
	chanTask := make(chan Task)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range chanTask {
				mu.Lock()
				errorsCountNow := errorsCount
				mu.Unlock()

				if maxErrorsCount > 0 && errorsCountNow < maxErrorsCount {
					if err := task(); err != nil {
						mu.Lock()
						errorsCount++
						mu.Unlock()
					}
				}
			}
		}()
	}

	go func() {
		defer close(chanTask)
		for _, task := range tasks {
			mu.Lock()
			errCountNow := errorsCount
			mu.Unlock()

			if errCountNow >= maxErrorsCount {
				break
			}

			chanTask <- task
		}
	}()

	wg.Wait()
	if maxErrorsCount <= 0 || errorsCount < maxErrorsCount {
		return nil
	}
	return ErrErrorsLimitExceeded
}
