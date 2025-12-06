package main

import "fmt"

func main(){
	frase := "Ana comeu a banana"
	tamanho := len(frase)
	fmt.Printf("\nO tamanho da frase é: %d\n", tamanho)
	primeiraPalavra := frase[:3]
	fmt.Printf("Primeira palavra: %s\n", primeiraPalavra)
	segundaPalavra := frase[4:9]
	fmt.Printf("Segunda palavra: %s\n", segundaPalavra)
	ultimaPalavra := frase[12:]
	fmt.Printf("Última palavra: %s\n", ultimaPalavra)
}
