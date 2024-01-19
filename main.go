package main

import "fmt"

// ? se usa coma y luego el tipo para saber que esos argumentos son de ese tipo
func dobleNumero(a,b int){
	fmt.Printf("%d y %d es igual a : %d\n",a,b,a + b)
}

// ? para que una funcion retorne 2 valores se especifica con otro ()

func doble(a int) (b,c int) {
	return a , a *4
}

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

	dobleNumero(5,8)
	dobleNumero(14,22)

	//? para las funciones con retorno de varios datos se guardan en numero de variables que retorna si en caso que no quiera algun dato lo especificamos con _ 


	value1 , value2 := doble(8)

	fmt.Println(value1,value2)

	value3 , _ := doble(13)

	fmt.Println("se quiere el primer valor",value3)

	_,value4 := doble(25)

	fmt.Println("se quiere el segundo valor",value4)

	// ? ciclo for igual que en javascript

	for i := 0; i < 10 ; i++{
		fmt.Println("ciclo",i)
	}

	// ? for while 

	counter := 0

	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	//? for forever este ciclo es infinito 
	// forever := 0
	// for{
	// 	fmt.Println(forever)
	// 	forever++
	// }

	

}
