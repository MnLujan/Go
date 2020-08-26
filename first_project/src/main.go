package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hola perri")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(nil)
	} else {
		fmt.Println(text)
	}
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
}
