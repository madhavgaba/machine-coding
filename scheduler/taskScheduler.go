package scheduler

import (
	"task-scheduler/entity"
	"task-scheduler/taskStore"
)

type TaskScheduler struct {
	taskQueue *taskStore.TaskQueue
}

func NewTaskScheduler(taskQueue *taskStore.TaskQueue) *TaskScheduler {
	return &TaskScheduler{
		taskQueue: taskQueue,
	}
}

func (t *TaskScheduler) Schedule(taskInfo entity.Task) {
	t.taskQueue.AddTask(entity.Task{
		Id:            taskInfo.Id,
		Type:          taskInfo.Type,
		Interval:      taskInfo.Interval,
		ExecutionTime: taskInfo.ExecutionTime,
		Runnable:      taskInfo.Runnable,
	})
}
