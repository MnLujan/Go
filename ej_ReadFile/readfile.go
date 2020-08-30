package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/* Abro el archivo */
	file, err := os.Open("./file.txt")

	if err != nil {
		fmt.Println("No se pudo abrir el archivo")
		return
	}

	/* Genero el scanner */
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i++
		linea := scanner.Text()
		fmt.Println("Numero de linea " + strconv.Itoa(i) + ": " + linea)
	}
}
