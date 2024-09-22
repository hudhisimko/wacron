package worker

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	maxWorker := 5
	wp := newWorkerPool(maxWorker)
	// Start the worker pool
	wp.Run()

	// Number of tasks to add
	numTasks := 100

	// WaitGroup to synchronize goroutines
	var wg sync.WaitGroup
	wg.Add(numTasks)

	// Add tasks concurrently
	for i := 0; i < numTasks; i++ {
		go func(i int) {
			wp.AddTask(&Task{
				ID: i,
				Run: func(int) {
					// Simulate some task execution
					time.Sleep(time.Millisecond * 10)
					wg.Done()
				},
			})
		}(i)
	}

	// Stop the worker pool concurrently
	go func() {
		<-time.After(time.Millisecond * 10) // Wait a bit before stopping
		<-wp.Stop()
	}()

	// Wait for all tasks to complete
	wg.Wait()
}

func TestWorkerPool_AddTask(t *testing.T) {
	// Create a new worker pool with a max worker of 3
	wp := newWorkerPool(5)
	wp.Run()

	// Create a wait group to wait for all tasks to finish
	var wg sync.WaitGroup

	// Define a function to execute tasks
	counter := atomic.Int32{}
	execFunc := func(workerID int) {
		// Simulate task execution time
		time.Sleep(50 * time.Millisecond)
		counter.Add(1)
	}

	// Add tasks with the same ID
	numTasks := 5
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			wp.AddTask(&Task{ID: id, Run: execFunc})
		}(i % 2)
	}
	const runningTasks = 2 // because (i%2) will be 0 or 1

	time.Sleep(10 * time.Millisecond)
	// Check if the number of tasks executed matches the number of tasks added
	wp.tasksInProgressLock.Lock()
	if len(wp.tasksInProgress) != runningTasks {
		t.Errorf("Expected %d tasks to be in progress, got %d", numTasks, len(wp.tasksInProgress))
	}
	wp.tasksInProgressLock.Unlock()

	// Wait for all tasks to finish
	wg.Wait()

	// Stop the worker pool and wait for it to finish
	<-wp.Stop()

	if counter.Load() != runningTasks {
		t.Errorf("Counter should be same as ran tasks, got %d, should be %d", counter.Load(), runningTasks)
	}

	// Check if the number of tasks in progress is zero after stopping the worker pool
	if len(wp.tasksInProgress) != 0 {
		t.Errorf("Expected 0 tasks to be in progress after stopping the worker pool, got %d", len(wp.tasksInProgress))
	}
}
