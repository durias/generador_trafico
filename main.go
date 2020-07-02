/*
  Cliente HTTP en Go con net/http
  Ejemplo de petición HTTP POST enviando datos JSON
  en Golang

*/
package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"bufio"
	"os"
)

type Caso struct {
	Edad         int
	Contagio     string
	Departamento string
	Nombre       string
	Estado       string
}

func enviarCaso(caso Caso) {

	clienteHttp := &http.Client{}
	// Url del servidor
	url := "http://localhost:8080"

	StructComoJson, err := json.Marshal(caso) //Se convierte el struct a json
	if err != nil {
		log.Fatalf("Error codificando usuario como JSON: %v", err)
	}
	//Se hace la peticion Post
	peticion, err := http.NewRequest("POST", url, bytes.NewBuffer(StructComoJson))
	if err != nil {
		log.Fatalf("Error creando petición: %v", err)
	}

	respuesta, err := clienteHttp.Do(peticion)
	if err != nil {
		log.Fatalf("Error haciendo petición: %v", err)
	}
	defer respuesta.Body.Close()

	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		log.Fatalf("Error leyendo respuesta: %v", err)
	}
	//Respuesta de la API
	respuestaString := string(cuerpoRespuesta)
	log.Printf("La api dice: '%s'", respuestaString)

}

func main() {
	texto_archivo := ""
	var casos []Caso
	file, err := os.Open("datos.fs")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		texto_archivo += scanner.Text()
	}

	json.Unmarshal([]byte(texto_archivo), &casos)
	//fmt.Printf("casos : %+v", casos)

	for i := 0; i < len(casos); i++ {
		enviarCaso(casos[i])
	}

}
