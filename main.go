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

	// ? Operaciones ?
	// * Suma
	suma := 2 + 3
	fmt.Println("la suma es: ",suma)
	// * Resta
	resta := 2 - 3
	fmt.Println("la resta es: ",resta)
	// * Multiplicacion
	multi := 2 * 3
	fmt.Println("la multi es: ",multi)
	// * Division
	division := 2 / 3
	fmt.Println("la division es: ",division)
	// * Modulo
	modulo := 10 % 2
	fmt.Println("El reciduo es ",modulo)
	// * Incremento
	modulo++
	fmt.Println("incremento ",modulo)
	// * Decremento
	modulo--
	fmt.Println("decremento",modulo)

	// ? funciones com fmt
	// * imprimir en la consola con salto de linea 
	fmt.Println("este es un mensaje con salto de linea")

	// * imprimir variables dentro de una string %s para string %d para entero %v si no se sabe el tipo de dato, para saber el tipo de dato se usa %T
	age := 25
	fmt.Printf("age: %T\n",age)
	fmt.Printf("%s tiene %d\n",nombre,age)
	fmt.Printf("%v tiene %v\n",nombre,age)
	// * Sprintf no imprime en consola solo guarda como string
	message := fmt.Sprintf("%s tiene %d",nombre,age)
	fmt.Println(message)

}
