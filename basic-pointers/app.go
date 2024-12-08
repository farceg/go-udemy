package main

import "fmt"

/*
* No usar pointers para ejemplos como los de este aplicación,
* pues un int ocupa muy poco espacio en memoria entonces
* no es necesario. Es más pesado crear el pointer que usar
* una copia del int.
 */

func main() {
	age := 27

	var agePointer = &age

	fmt.Println("Dirección inicial:", agePointer)
	fmt.Println("")
	fmt.Println("")
	test1(agePointer)
	fmt.Println("-----------------------------")
	test2(age)
}

// Conserva la misma dirección en memoria
func test1(age *int) {
	fmt.Println("Dirección 1:", age)
	fmt.Println("Valor variable 1:", *age)
	*age = 10 // Mutar el valor del puntero para evitar un retorno
}

// Conserva distinta dirección en memoria
func test2(age int) {
	fmt.Println("Dirección 2:", &age)
	fmt.Println("Valor variable 2:", age)
}
