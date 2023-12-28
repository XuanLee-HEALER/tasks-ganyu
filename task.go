package tasksganyu

type Task interface {
}

type BatchTask interface {
}

type TaskWrapper struct {
	task Task
}

type BatchTaskWrapper struct {
	task Task
}
