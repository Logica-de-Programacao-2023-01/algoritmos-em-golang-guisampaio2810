package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Qual o numero? ")
	fmt.Scan(&numero)

	fmt.Println("O seu numero vezes 2 é? ", numero*2)
	fmt.Println("O seu numero vezes 3 é? ", numero*3)
	fmt.Println("O seu numero vezes 4 é? ", numero*4)
}
