package main

import (
	"fmt"
	"log"
	"strconv"
)

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
	//? condicionales
	x := 5
	y := 10
	
	if x == 5 && y != 6 {
		fmt.Println("condicional con and && ")
	}

	if x == 3 || y == 10 {
		fmt.Println("condicional con or ||")
	}

	//? conversion de texto a numero
	valor, err := strconv.Atoi("53")

	//? nil se usa para saber si una funcion no tiene error
	// * el log lanzara en consola el error 

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(valor)
	// ? ejercicio para saber si un numero dado es par o impar

	result := 5

	if result % 2 == 0 {
		fmt.Println("es par")
	}else{
		fmt.Println("es impar")
	}

	// ? condicion swicht

	switch modulo := 7 % 2 ; modulo{
	case 0:
		fmt.Println("es par")
	default:
		fmt.Println("es impar")
	}

	valor2 := -400
	switch{
	case valor2 > 100:
		fmt.Println("es mayor a 100")
	case valor2 < 0:
		fmt.Println("es menor a cero")
	default:
		fmt.Println("no entro a ninguna condicion")
	}

	// ? keyword defer se usa para ejecutar algo antes de que muera una funcion es bueno usarlo para cerrar conexion con DB o para cerrar archivos

	defer fmt.Println("ejecuto antes de que muera ")

	// ? otras keyword son el continue y break para los ciclos como en otros lenguajes 
	for i := 0 ; i < 10 ; i++ {
		if i == 3{
			fmt.Println("es 3")
			continue
		}
		if i == 9{
			break
		}
		fmt.Println(i)
	}

	// ? arrays con go son inmutables en tamaño pero se pueden modificar sus elementos

	var array [4]int
	array[0] = 2
	array[1] = 200
	fmt.Println(array,len(array),cap(array))

	//? slice es un tipo de array pero a diferencia en este no se especifica su longitud

	slice := []int {1,2,3,4,5,6}
	fmt.Println(slice,len(slice),cap(slice))

	// * metodos de slice al igual que en python se usa los slicing para cortar diferentes partes
	fmt.Println(slice[0])
	fmt.Println(slice[:5])
	fmt.Println(slice[2:4])
	fmt.Println(slice[:3])
	// * append se debe especificar a que slice se va agegar el dato requerido
	slice = append(slice,10)
	fmt.Println(slice)
	// * en esta parte podemos agregar un slice a otro slice automaticamente pone los 3 puntos lo que hace es un destructuring
	newSlice := []int {22,23,24,25}
	slice = append(slice, newSlice...)

	fmt.Println(slice)

	// * recorrido de slice en caso de no querer el iterador i se hace igual que en las funciones de retorno de varios valores con guion al piso
	for i , valor := range slice{
		fmt.Println(i,valor)
	}

	for _ , valor := range slice{
		fmt.Println(valor)
	}

	// * en caso de solo querer el iterador solo se pone el valor de iterador sin necesidad de escapar el valor

	for i := range slice {
		fmt.Println(i)
	}

}
