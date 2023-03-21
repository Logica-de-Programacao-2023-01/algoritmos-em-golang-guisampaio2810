package main

import "fmt"

func main() {
	var peso float64
	var altura float64
	fmt.Print("Seu peso é:")
	fmt.Scan(&peso)
	fmt.Println("sua altura é")
	fmt.Scan(&altura)
	fmt.Println("Seu IMC é:", peso/(altura*altura))

}
