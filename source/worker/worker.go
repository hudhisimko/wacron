package worker

import "sync"

type Task struct {
	ID  int
	Run func(workerID int)
}

type Worker struct {
	*workerPool
}

func NewWorker(maxWorker int) *Worker {
	return &Worker{
		workerPool: newWorkerPool(maxWorker),
	}
}

type workerPool struct {
	maxWorker   int
	queuedTaskC chan *Task

	tasksInProgress     map[int]struct{}
	tasksInProgressLock sync.Mutex

	finish      chan struct{}
	stopped     bool // this prevent sending task if worker has been stopped
	stoppedLock sync.Mutex

	running bool
}

func newWorkerPool(maxWorker int) *workerPool {
	return &workerPool{
		maxWorker:       maxWorker,
		queuedTaskC:     make(chan *Task),
		finish:          make(chan struct{}),
		tasksInProgress: map[int]struct{}{},
	}
}

func (wp *workerPool) Run() {
	// ensure run executed once
	if wp.running {
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(wp.maxWorker)
	for i := 0; i < wp.maxWorker; i++ {
		go func(workerID int) {
			for task := range wp.queuedTaskC {
				// execute the task
				task.Run(workerID)

				wp.tasksInProgressLock.Lock()
				delete(wp.tasksInProgress, task.ID)
				wp.tasksInProgressLock.Unlock()
			}

			wg.Done()
		}(i)
	}

	// wait until task finish before stopped
	go func() {
		wg.Wait()
		close(wp.finish)
	}()

	wp.running = true
}

func (wp *workerPool) AddTask(task *Task) {
	wp.stoppedLock.Lock()
	defer wp.stoppedLock.Unlock()

	if wp.stopped {
		return
	}

	// these lock ensure same id won't be executed multiple times
	wp.tasksInProgressLock.Lock()
	if _, ok := wp.tasksInProgress[task.ID]; ok {
		wp.tasksInProgressLock.Unlock()
		return
	}
	wp.tasksInProgress[task.ID] = struct{}{}
	wp.tasksInProgressLock.Unlock()

	wp.queuedTaskC <- task
}

func (wp *workerPool) Stop() <-chan struct{} {
	wp.stoppedLock.Lock()
	wp.stopped = true
	wp.stoppedLock.Unlock()

	close(wp.queuedTaskC)
	return wp.finish
}
