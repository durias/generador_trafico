package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Caso struct {
	Edad         int
	Contagio     string
	Departamento string
	Nombre       string
	Estado       string
}

func recibirPost(w http.ResponseWriter, r *http.Request) {
	// Variable para guardar la info de entrada
	var info Caso

	// Aqui decodifica la info para convertirla en Json, si no funciona da error xd
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	/*Si la ejecucion llega a esta linea es que el struct llego bien y se puede usar
	Para lo que sea necesario */
	fmt.Fprintf(w, "Recibi caso: %+v", info)
}

func main() {
	http.HandleFunc("/", recibirPost)
	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
