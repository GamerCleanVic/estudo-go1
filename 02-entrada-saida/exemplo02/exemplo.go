package main

import "fmt"

func main(){
	fmt.Println("\nDigite o seu nome:")
	var nome string
	fmt.Scanf("%s", &nome)

	serie := "Lógica de Programação em Go"
	fmt.Printf("Seja bem bindo à série %s, %s!\n\n", serie, nome)
}
