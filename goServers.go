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
	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, servidor := range servidores {
		go revisarServidor(servidor, chanel)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-chanel)
	}
	tiempoPasado := time.Since(inicio)
	fmt.Printf("Done in %s", tiempoPasado)
}
