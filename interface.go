package main

import "fmt"

type animal interface {
	move() string
}

type dog struct{}
type fish struct{}
type bird struct{}

func (dog) move() string {
	return "I'm a dog and i walk"
}

func (fish) move() string {
	return "I'm a fish and i am swiming"
}

func (bird) move() string {
	return "I'm a fish and i am flying"
}

func moveAnimal(a animal) {
	fmt.Println((a.move()))
}

func main() {
	p := dog{}
	moveAnimal(p)
	pe := fish{}
	moveAnimal(pe)
	pa := bird{}
	moveAnimal(pa)
}
