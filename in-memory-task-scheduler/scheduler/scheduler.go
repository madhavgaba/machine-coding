package scheduler

import "task-scheduler/entity"

type Scheduler interface {
	Schedule(taskInfo entity.Task)
}
