package main

import "fmt"

func main() {
	var precoProduto float64
	fmt.Print("Qual preço do produto?")
	fmt.Scan(&precoProduto)
	fmt.Println("O preço do produto com 10% é de:" precoProduto * 0.9)

}
