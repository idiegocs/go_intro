package variables

import "fmt"

func ExecuteVariables() {

	fmt.Println("Init Variables-----------------------")

	var nombre string
	var edad int

	fmt.Println("defecto nombre:", nombre)
	fmt.Println("defecto Edad:", edad)

	nombre = "Juan"
	edad = 30

	fmt.Println("nombre:", nombre)
	fmt.Println("Edad:", edad)

	apellido := "Sanchez"
	fmt.Println("Apellido:", apellido)

	fmt.Printf("NombreCompleto: %s %s Edad: %d \n", nombre, apellido, edad)

	isOk := true
	fmt.Println("isOk:", isOk)

	const pi = 3.1416
	fmt.Println("Pi:", pi)

	fmt.Println("Fin Variables-----------------------")

}
