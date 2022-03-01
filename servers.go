package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()
	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, servidor := range servidores {
		revisarServidor(servidor)
	}

	tiempoPasado := time.Since(inicio)
	fmt.Printf("Done in %s", tiempoPasado)
}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)

	if err != nil {
		fmt.Println(servidor, "no esta disponible")
	} else {
		fmt.Println(servidor, "funciona correctamente")
	}

}
