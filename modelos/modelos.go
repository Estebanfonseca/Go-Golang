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


// ! las funciones y/o structs definidos en otros archivos su nombre debe empezar por mayuscula para que se puedan acceder a ellas desde otros archivos
type Pc struct{
	Marca string
	Ram int
	Disco int
}

func (myPC *Pc) DuplicarRam() {
	myPC.Ram *= 2
}

//? customizar la salida del print de la linea 305 de myPc

func (myPC Pc) String() string{
	return fmt.Sprintf("tengo %d de Ram , %d de disco y soy un %s",myPC.Ram,myPC.Disco,myPC.Marca)
}

// ? uso de interfaces y su aplicacion

// * esta interfase usa la funcion en comun para ejecutarse dependiendo de la struct a la que pertenece
type Figuras2D interface{
	Area() float64
}

// * aqui se crean los 2 struct 
type Cuadrado struct{
	Base float64
}

type Rectangulo struct{
	Base float64
	Altura float64
}


// * aqui creamos  las 2 funciones que se van ejecuntando con el mismo nombre pero usan struct diferente
func (c Cuadrado) Area() float64{
	return c.Base * c.Base
}

func (r Rectangulo) Area() float64{
	return r.Altura * r.Base
}

//* esta funcion usa la interfase creada y ejecuta su funcion 
func Calcular(f Figuras2D){
	fmt.Println("Area;",f.Area())
}
