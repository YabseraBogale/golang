package worker

import (
	"fmt"

	"github.com/YabseraBogal/golang/orchestrator/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("I Will Collect Stats")
}
func (w *Worker) RunTask() {
	fmt.Println("I start or stop a tasks")
}
func (w *Worker) StartTask() {
	fmt.Println("I will start task")
}
func (w *Worker) StopTask() {
	fmt.Println("I will stop task")
}
