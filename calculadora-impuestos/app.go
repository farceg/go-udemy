package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var ingresos float64
	var gastos float64
	var impuestos float64

	err := getUserInput("Ingresos: ", &ingresos)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = getUserInput("Gastos: ", &gastos)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = getUserInput("Impuestos: ", &impuestos)
	if err != nil {
		fmt.Println(err)
		return
	}

	ganancias, gananciasMenosImpuestos, ratio := calculateValues(ingresos, gastos, impuestos)

	storeValues(ganancias, gananciasMenosImpuestos, ratio)

	fmt.Println("-------------------------------------")
	fmt.Printf("Ganancias: %.2f\nGanancias menos impuestos: %.2f\nRatio: %.2f", ganancias, gananciasMenosImpuestos, ratio)

}

func getUserInput(info string, value *float64) error {
	fmt.Print(info)
	fmt.Scan(value)
	if *value <= 0 {
		return errors.New("el valor ingresado no puede ser cero o negativo")
	}
	return nil
}

func calculateValues(ingresos, gastos, impuestos float64) (ganancias, gananciasMenosImpuestos, ratio float64) {
	ganancias = ingresos - gastos
	gananciasMenosImpuestos = ganancias * (1 - impuestos/100)
	ratio = ganancias / gananciasMenosImpuestos
	return
}

func storeValues(ganancias, gananciasMenosImpuestos, ratio float64) {
	texto := fmt.Sprintf("Ganancias %.2f, ganancias menos impuestos %.2f, ratio %.2f", ganancias, gananciasMenosImpuestos, ratio)
	os.WriteFile("Resultado.txt", []byte(texto), 0644)
}
