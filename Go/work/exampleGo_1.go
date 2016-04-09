package main

import "fmt"

func main() {
  fmt.Print("Ingrese un numero: ")
  var input string
  fmt.Scanf("%s", &input)

  output := input

  fmt.Println(output)
}


/*EJERCICIO: Cambie el anterior programa para que en vez de capturar
	un decimal, capture un texto y lo imprima.*/