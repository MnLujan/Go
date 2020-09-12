package main

import (
	"./dummy"
	"fmt"
)

func main() {
	init := dummy.Test
	fmt.Println("Podemos ver que la funcion del packete Dummy devuelve: " + dummy.Dummy())
	fmt.Println("Y su variable publica es: " + init)
}
