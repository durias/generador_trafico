/*
  Cliente HTTP en Go con net/http
  Ejemplo de petición HTTP POST enviando datos JSON
  en Golang

*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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

func main() {

	var total_datos int
	var total_hilos int
	var entrada string
	var path string
	var array_casos []Caso

	//var flag_errors = true
	fmt.Println("*******************************************************************************")
	fmt.Println("*                    Bienvenido al sistema  DOCH 2020                         *")
	fmt.Println("*                                                                             *")
	fmt.Println("*******************************************************************************")

	for {
		fmt.Print("Digite la cantidad de datos a cargar: ")
		fmt.Scanf("%s\n", &entrada)
		n, _ := strconv.Atoi(entrada)
		total_datos = n
		if total_datos > 0 {
			break
		} else {
			fmt.Println("Error: El dato debe ser entero mayor a cero.")
		}
	}

	for {
		fmt.Print("Digite la cantidad de hilos: ")
		fmt.Scanf("%s\n", &entrada)
		n, _ := strconv.Atoi(entrada)
		total_hilos = n
		if total_hilos > 0 {
			break
		} else {
			fmt.Println("Error: El dato debe ser entero mayor a cero.")
		}
	}

	for {
		fmt.Print("Escriba la url del archivo: ")
		fmt.Scanf("%s\n", &path)
		fmt.Printf("Intentando leer datos desde la ruta: %s ... \n", path)
		fmt.Println("-------------------------------------------------------------------------------")
		array_casos = get_array_casos(path)
		if array_casos != nil {
			break
		} else {
			fmt.Println("Error: la ruta es incorrecta.")
		}
		break
	}

	fmt.Println("Información completada con éxito, presione ENTER para iniciar concurrencia.")
	fmt.Scanln()
	iniciar_concurrencia(total_datos, total_hilos, array_casos)
	fmt.Println("\nCasos enviados con exito presione ENTER para salir.")
	fmt.Scanln()

}

func get_array_casos(ruta string) []Caso {
	texto_archivo := ""
	var casos []Caso
	file, err := os.Open(ruta)
	if err != nil {
		return nil
		//log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		texto_archivo += scanner.Text()
	}
	json.Unmarshal([]byte(texto_archivo), &casos)
	return casos
}

func iniciar_concurrencia(total int, hilos int, arreglo []Caso) {

	var longitud_archivo int
	var contador int
	contador = 0
	longitud_archivo = len(arreglo)

	fmt.Printf("Longitud del archivo: %v ... \n", longitud_archivo)
	fmt.Printf("Datos por hilo: %v ... \n", total)

	for cont_hilos := 0; cont_hilos < hilos; cont_hilos++ {
		var pack_jsons []Caso
		for i := 0; i < total; i++ {
			if contador >= longitud_archivo {
				contador = 0
			}
			pack_jsons = append(pack_jsons, arreglo[contador])
			//fmt.Printf("%d\n", contador)
			contador += 1
		}
		// llamada al metodo enviar paquete en una nueva rutina
		go enviar_paquete(pack_jsons)
	}
}

func enviar_paquete(arreglo []Caso) {

	fmt.Printf("valor de i: %v ... \n", len(arreglo))
	for i := 0; i < len(arreglo); i++ {
		//time.Sleep(25 * time.Millisecond)
		enviarCaso(arreglo[i])
	}

}

func enviarCaso(caso Caso) {

	clienteHttp := &http.Client{}
	// Url del servidor
	//url := "http://34.68.15.208:3001"
	url := "http://localhost:3001"

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
	fmt.Printf("La api dice: '%s' \n", respuestaString)

}
