package main

import "fmt"

func main(){
	frase := "Olá, mundo!"
	primeiroCaracter := frase[0]
	fmt.Printf("Primeiro caracter: %c\n", primeiroCaracter)
	indiceUltimoCaracter := len(frase) -1
	ultimoCaracter := frase[indiceUltimoCaracter]
	fmt.Printf("Último caracter: %c\n", ultimoCaracter)
}
