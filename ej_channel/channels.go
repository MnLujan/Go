package main

import "fmt"

func main() {
	/* Creo el canal */
	channel := make(chan string)

	/* Se crea un codigo que concurrente */
	go func(channel chan string) {
		for {
			var buffer string
			/* Leo del teclado */
			fmt.Scanln(&buffer)
			channel <- buffer
		}
	}(channel)

	/* Espero a que haya una carga de datos */
	for {
		/* Leo del canal */
		msg := <-channel
		if msg == "exit" {
			fmt.Println("bye bye")
			break
		}
		fmt.Println("Dato leido desdel el canal: " + msg)
	}

}
