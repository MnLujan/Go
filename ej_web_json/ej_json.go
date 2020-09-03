package main

import (
	"encoding/json"
	"net/http"
)

type Data struct {
	Name string
	Edad int
}

type Juan struct {
	Apellido string `json:"apellido:"`
	Edad     int    `json:"Esta viejo?: "`
}

func main() {
	http.HandleFunc("/juan", func(writer http.ResponseWriter, request *http.Request) {
		juancito := Juan{"Gismondi", 31}

		/* Recordar que con el operador "_" estoy diciendo q no me interesa el valor que retorna */
		_ = json.NewEncoder(writer).Encode(juancito)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Data{"Martin", 23}
		_ = json.NewEncoder(w).Encode(data)
	})

	_ = http.ListenAndServe(":8000", nil)
}
