package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(servidor string, c chan string) {
	_, err := http.Get(servidor)

	if err != nil {
		c <- servidor + " no esta disponible"
	} else {
		c <- servidor + " funciona correctamente"
	}

}

func main() {
	chanel := make(chan string)
	inicio := time.Now()
	i := 0

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for i < 2 {

		for _, servidor := range servidores {
			go revisarServidor(servidor, chanel)
		}

		time.Sleep(1 * time.Second)
		fmt.Println(<-chanel)

		i++
	}

	tiempoPasado := time.Since(inicio)
	fmt.Printf("Done in %s", tiempoPasado)
}
