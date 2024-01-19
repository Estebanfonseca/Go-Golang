package modelos

import "fmt"

// * para hacer que sea acceso publico la clase debe empezar con mayuscula al igual que sus valores
type CarPublic struct{
	Anio int
	Modelo string
}

// * al igual que la clase las funciones funcionan igual para acceder desde cualquier lado el nombre debe ser mayuscula 

func Texto(){
	fmt.Println("hola")
}