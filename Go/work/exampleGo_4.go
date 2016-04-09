package main

import "fmt"

func main() {
	x := []int{
	  48,96,86,68,
	  57,82,63,70,
	  37,34,83,27,
	  19,97, 9,17,
	}
	var menor = encontrarMenor(x)
	
	fmt.Println(menor)
}
func encontrarMenor(arreglo[] int) int{

	var menor = arreglo[0]

	for i := 0; i < len(arreglo); i++{
		if arreglo[i] < menor{
			menor = arreglo[i]
		}
	}
	return  menor
}