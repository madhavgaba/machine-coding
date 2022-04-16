package main

import (
	"fmt"
	"sync"
	"task-scheduler/constants"
	"task-scheduler/entity"
	"task-scheduler/internal/priority_queue"
	"task-scheduler/scheduler"
	"task-scheduler/taskStore"
	"time"
)

func getTasks() []entity.Task {
	return []entity.Task{
		{
			Id:            1,
			Type:          "once",
			Interval:      0,
			ExecutionTime: time.Now().Unix() + 15,
			Runnable: func() {
				fmt.Println("task #1, ONCE, executed at", time.Now().Unix())
			},
		},
		{
			Id:            2,
			Type:          "recurring",
			Interval:      10,
			ExecutionTime: time.Now().Unix() + 7,
			Runnable: func() {
				fmt.Println("task #2, RECURRING, executed at", time.Now().Unix())
			},
		},
		{
			Id:            3,
			Type:          "recurring",
			Interval:      6,
			ExecutionTime: time.Now().Unix() + 10,
			Runnable: func() {
				fmt.Println("task #3, RECURRING, executed at", time.Now().Unix())
			},
		},
	}
}

func main() {

	tasks := getTasks()
	pq := priority_queue.NewPriorityQueue()
	var wg sync.WaitGroup

	tQueue := taskStore.NewTaskQueue(constants.MAX_WORKERS, pq, &wg)
	taskScheduler := scheduler.NewTaskScheduler(tQueue)

	for i := 0; i < len(tasks); i++ {
		taskScheduler.Schedule(tasks[i])
	}

	wg.Wait()
}
