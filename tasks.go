package main

import "fmt"

type taskList struct {
	tasks []task
}

func (t *taskList) addToList(tl task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) removeFromList(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

type task struct {
	name        string
	description string
	complete    bool
}

func (t *task) markComplete() {
	t.complete = true
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func (t *task) updateName(name string) {
	t.name = name
}
func main() {
	t := task{
		name:        "Completar curso de golangooooo",
		description: "Completar el curso esta semana",
		complete:    false,
	}
	t1 := task{
		name:        "Completar curso de python",
		description: "Completar el curso esta semana",
		complete:    false,
	}
	t2 := task{
		name:        "Completar curso de go",
		description: "Completar el curso esta semana",
		complete:    false,
	}

	list := taskList{
		tasks: []task{t1, t2},
	}
	fmt.Println(list.tasks[0])
	list.addToList(t)
	fmt.Println(len(list.tasks))
	list.removeFromList(1)
	fmt.Println(len(list.tasks))
}
