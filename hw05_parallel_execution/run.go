package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	maxErrors := int32(m)
	countErrors := int32(0)

	ch := make(chan Task)
	wg := sync.WaitGroup{}

	// Create worker goroutines.
	for w := 0; w < n; w++ {
		wg.Add(1)
		go worker(ch, &wg, &countErrors)
	}

	// Send tasks to workers via channel.
	for _, task := range tasks {
		// функция должна останавливать свою работу, если произошло m ошибок.
		// m <= 0 -- игнорировать ошибки.
		if isMaxErr(atomic.LoadInt32(&countErrors), maxErrors) {
			break
		}
		ch <- task
	}
	close(ch)

	wg.Wait()

	if isMaxErr(countErrors, maxErrors) {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func isMaxErr(countErrors int32, maxErrors int32) bool {
	return countErrors >= maxErrors && maxErrors > 0
}

func worker(ch chan Task, wg *sync.WaitGroup, countErrors *int32) {
	defer wg.Done()

	for task := range ch {
		if err := task(); err != nil {
			atomic.AddInt32(countErrors, 1)
		}
	}
}
