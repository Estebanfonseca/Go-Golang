package main

import "fmt"

func main() {
	fmt.Println("hola mundo")

	// ? Declaracion de variables

	// * se crea la variable y se asigna autmaticamente el tipo de dato
	nombre := "esteban"

	// * se crea la varialbe y se asigna el tipo de dato
	var apellido string = "gutierrez"

	// * se crea la variable con el tipo pero no el valor 
	var edad int

	// * si no se usa la  variable el editor lo marca como error
	fmt.Println(nombre, apellido, edad)

	// ? Tipos de datos y su valor por defecto
	
	// * int por defecto es cero
	var a int
	// * float por defecto es cero 
	var b float64
	// * string por defecto es un string vacio
	var c string
	// * bool por defecto es false
	var d bool

	fmt.Println(a,b,c,d)
}
