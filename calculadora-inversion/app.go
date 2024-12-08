package main

import (
	"fmt"
	"math"
)

func main() {

	const inflacion float64 = 2.5

	var valorInversion float64
	var anios float64
	var valorRetorno float64

	fmt.Print("Ingresa el valor a invertir: ")
	fmt.Scan(&valorInversion)
	fmt.Print("Ingresa el periodo de tiempo en años: ")
	fmt.Scan(&anios)
	fmt.Print("Ingresa el retorno esperado: ")
	fmt.Scan(&valorRetorno)

	var valorFuturo float64 = valorInversion * math.Pow(1+valorRetorno/100, anios)
	var valorFuturoReal float64 = valorFuturo / math.Pow(1+inflacion/100, anios)

	fmt.Println("-----------------------------------")
	fmt.Printf("Valor sin inflacion: %.2f\n", valorFuturo)
	fmt.Printf("Valor con inflacion: %.2f", valorFuturoReal)

}
