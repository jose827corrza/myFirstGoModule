package utils

import "fmt"

// Recordar que una funcion que se quiere exportar, DEBE iniciar con mayuscula
// si se inicia con minuscula seria una funcion privada,
// que solo podria ser llamada dentrode este modulo

func HelloWorld() {
	fmt.Println("Hello from my first Go module!!")
}
