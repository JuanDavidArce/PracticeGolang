package main

import "fmt"

type dog struct{}
type fish struct{}
type bird struct{}

func (dog) walk() string {
	return "I'm a dog and i walk"
}

func (fish) swim() string {
	return "I'm a fish and i am swiming"
}

func (bird) fly() string {
	return "I'm a fish and i am flying"
}

func moveDog(d dog) {
	fmt.Println((d.walk()))
}
func moveFish(f fish) {
	fmt.Println((f.swim()))
}
func moveBird(b bird) {
	fmt.Println((b.fly()))
}

func main() {
	p := dog{}
	moveDog(p)
	pe := fish{}
	moveFish(pe)
	pa := bird{}
	moveBird(pa)
}
