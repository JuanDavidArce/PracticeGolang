package main

import "fmt"

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
		name:        "Completar curso de go",
		description: "Completar el curso esta semana",
		complete:    false,
	}
	fmt.Printf("%+v\n", t)
	t.markComplete()
	t.updateName("FInalizar curso de go")
	t.updateDescription("Terminar cuato antes")
	fmt.Printf("%+v\n", t)
}
