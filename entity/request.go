package entity

type Task struct {
	Index         int
	Id            int
	Type          string
	ExecutionTime int64
	Interval      int64
	Runnable      func()
}
