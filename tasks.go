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

func (t *taskList) printList() {
	for _, task := range t.tasks {
		fmt.Println("Name", task.name)
		fmt.Println("Description", task.description)
	}
}

func (t *taskList) printListComplete() {
	for _, task := range t.tasks {
		if task.complete {
			fmt.Println("Name", task.name)
			fmt.Println("Description", task.description)
		}

	}
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
	t3 := task{
		name:        "Completar curso de node",
		description: "Completar el curso esta semana",
		complete:    false,
	}

	list := taskList{
		tasks: []task{t1, t2},
	}
	list.addToList(t3)
	list.printList()
	list.tasks[0].complete = true
	fmt.Println("Task complete")
	list.printListComplete()

	mapTasks := make(map[string]*taskList)

	mapTasks["Juan David"] = &list
	fmt.Println(mapTasks)

}
