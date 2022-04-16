package taskStore

import (
	"sync"
	"task-scheduler/constants"
	"task-scheduler/entity"
	"task-scheduler/internal/priority_queue"
	"time"
)

type TaskQueue struct {
	Queue *priority_queue.PriorityQueue
	Lock  *sync.RWMutex
}

func (t *TaskQueue) AddTask(task entity.Task) {
	t.Lock.Lock()
	t.Queue.Insert(task)
	t.Lock.Unlock()
}

func (t *TaskQueue) Execute() {
	for {
		t.Lock.Lock()

		if t.Queue.Len() > 0 {
			task := t.Queue.Poll()
			now := time.Now().Unix()

			diff := task.ExecutionTime - now

			if diff > 0 {
				time.Sleep(time.Duration(diff) * time.Second)
			}

			task.Runnable()

			if task.Type == constants.RECURRING {
				nextScheduledTime := time.Now().Unix() + task.Interval
				task.ExecutionTime = nextScheduledTime
				t.Queue.Insert(task)
			}
		}

		t.Lock.Unlock()
	}
}

func NewTaskQueue(workers int, pq *priority_queue.PriorityQueue, wg *sync.WaitGroup) *TaskQueue {
	lock := sync.RWMutex{}

	taskQueue := TaskQueue{
		Queue: pq,
		Lock:  &lock,
	}

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			taskQueue.Execute()
		}()
	}

	return &taskQueue
}
