package priority_queue

import (
	"container/heap"
	"task-scheduler/entity"
)

type PriorityQueue []entity.Task

var taskPQ PriorityQueue

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].ExecutionTime < pq[j].ExecutionTime
}

func (pq *PriorityQueue) Push(task interface{}) {
	newTask := task.(entity.Task)
	newTask.Index = pq.Len()
	*pq = append(*pq, newTask)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	if n == 0 {
		return nil
	}
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq *PriorityQueue) Insert(task entity.Task) {
	heap.Push(&taskPQ, task)
}

func (pq *PriorityQueue) Poll() entity.Task {
	return heap.Pop(&taskPQ).(entity.Task)
}

func (pq *PriorityQueue) GetInstance() *PriorityQueue {
	return &taskPQ
}

func NewPriorityQueue() *PriorityQueue {
	taskPQ = make(PriorityQueue, 0)
	heap.Init(&taskPQ)

	return &taskPQ
}
