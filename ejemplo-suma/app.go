package main

import (
	"fmt"
	// Se importa usando el modulo que se crea con go mod init <path>
	"ejemplo-suma.com/operations"
)

func main() {
	a, b := 55, 25
	var sum int = operations.Sumar(a, b)
	var resta int = operations.Restar(a, b)
	fmt.Printf("Resultado de sumar %v + %v: %v", a, b, sum)
	fmt.Printf("\nResultado de restar %v - %v: %v", a, b, resta)
}
