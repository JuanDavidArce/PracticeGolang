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
	fmt.Println(list.tasks[0])
	list.addToList(t3)

	// for i := 0; i < len(list.tasks); i++ {
	// 	fmt.Println("Index", i, "nombre", list.tasks[i].name)
	// }

	// for index, value := range list.tasks {
	// 	fmt.Println("Index", index, "nombre", value.name)
	// }
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
