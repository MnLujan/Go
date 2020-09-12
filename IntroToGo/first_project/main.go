package main

import (
	"fmt"
	"time"
)

type persona struct {
	edad     int
	apellido string
}

type User interface {
	hola() string
	esmayor() bool
}

func (this persona) esmayor() bool {
	if this.edad > 20 {
		return true
	} else {
		return false
	}
}

func (this persona) hola() string {
	return "Hola Forro"
}

func auth(user User) bool {
	return user.esmayor()
}

func main() {
	go func() {
		martin := new(persona)
		martin.apellido = "Lujan"
		martin.edad = 23
		fmt.Println(*martin)
		martin.edad++
		fmt.Println(*martin)
		if auth(martin){
			fmt.Println("Es mayor el pibe")
		}
	}()
	fmt.Println("Hola perri")
	arreglo := [5]int{1, 2, 3, 4}
	fmt.Println(arreglo)
	aux := arreglo[0:2]
	fmt.Println(aux)
	fmt.Println(arreglo)
	aux = append(aux, 5)
	fmt.Println(aux)
	aux2 := make([]int, 3, 6)
	fmt.Println(aux2)
	aux2 = append(aux2, 8)
	fmt.Println(aux2)
	copy(aux, aux2)
	fmt.Println(aux)
	time.Sleep(1000*time.Millisecond)
}
