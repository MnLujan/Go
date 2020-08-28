package main

import (
	"fmt"
)

type persona struct {
	edad int
	apellido string
	next *persona

}

func main() {
	fmt.Println("Hola perri")
	arreglo := [5]int{1,2,3,4}
	fmt.Println(arreglo)
	aux := arreglo[0:2]
	fmt.Println(aux)
	fmt.Println(arreglo)
	aux = append(aux,5)
	fmt.Println(aux)
	aux2 := make([]int,3,6)
	fmt.Println(aux2)
	aux2 = append(aux2, 8)
	fmt.Println(aux2)
	copy(aux, aux2)
	fmt.Println(aux)
	martin := new(persona)
	martin.apellido = "Lujan"
	martin.edad = 23
	fmt.Println(*martin)
	martin.edad++
	fmt.Println(*martin)
	
}
