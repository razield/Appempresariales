package main

import "fmt"

type Persona struct{
	estatura float64
	nombre string
}

func (persona *Persona) correr() string{
	return "corriendo"
}

func main() {
	persona:= Persona{1.70, "Juan"}
	fmt.Println(persona.nombre, persona.correr())
}

