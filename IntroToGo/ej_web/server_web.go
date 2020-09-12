package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	/* La ruta distinta de la raiz (/) debe ir primero */
	http.HandleFunc("/compra", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Numero de tarjeta") //Notese q la funcion se define y construye en el parametro mismo.
		//Una manera de simplificar codigo
	})
	/* Ubicacion especifica */
	http.HandleFunc("/archivo", leer)

	/* Buscara el archivo dentro de la raiz donde este el binario */
	http.HandleFunc("/otro/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path[6:])
		http.ServeFile(w, r, r.URL.Path[6:])
	})
	http.HandleFunc("/", recep)
	http.ListenAndServe(":8000", nil)
}

func recep(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome")
}

func leer(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
