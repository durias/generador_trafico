package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//var cont int

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
	var json_cadena string

	// Aqui decodifica la info para convertirla en Json, si no funciona da error xd
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	/*Si la ejecucion llega a esta linea es que el struct llego bien y se puede usar
	Para lo que sea necesario */

	json_cadena = "{" + "\"Edad\":" + strconv.Itoa(info.Edad) + "," + "\"Contagio\":\"" + info.Contagio + "\"," + "\"Departamento\":\"" + info.Departamento + "\"," + "\"Nombre\":\"" + info.Nombre + "\"," + "\"Estado\":\"" + info.Estado + "\"" + "}"
	fmt.Fprintf(w, "Recibi caso: %+v", json_cadena) //Enviar respuesta al cliente
	fmt.Printf("%v \n", json_cadena)                //Imprimir en la consola del server
}

func main() {
	http.HandleFunc("/", recibirPost)
	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
