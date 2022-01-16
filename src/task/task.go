package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) addTaskList(tl *task) {
	t.tasks = append(t.tasks, tl)
}

type task struct {
	name        string
	description string
	completed   bool
}

func (t task) completeTask() {
	t.completed = true
}
func main() {
	t := &task{
		name:        "task",
		description: "description task",
		completed:   false,
	}
	t2 := &task{
		name:        "task 2",
		description: "description task 2",
		completed:   false,
	}
	t3 := &task{
		name:        "task 3",
		description: "description task 3",
		completed:   false,
	}
	list := &taskList{
		tasks: []*task{
			t, t2,
		},
	}
	list.addTaskList(t3)
	fmt.Println(len(list.tasks))
}
