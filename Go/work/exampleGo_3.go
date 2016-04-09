package main

import "fmt"

func main() {
	x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 9,17,
}
	var menor int = x[0]
	for i := 0; i < len(x); i++{
		if x[i] < menor{
			menor = x[i]
		}
	}
	
	fmt.Println(menor)
}

/*EJERCICIO: Escribir un programa que encuentre el numero menor en el 
siguiente arreglo y lo imprima:

x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 9,17,
}*/