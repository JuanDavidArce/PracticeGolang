package main

import "fmt"

func changeValue(a *int) {
	*a = 36
}

func main() {
	x := 25
	fmt.Println(x)
	changeValue(&x)
	fmt.Println(x)

}
