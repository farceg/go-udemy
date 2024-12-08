package main

import (
	"fmt"

	"example.com/ejemplo-banco/fileops" // Importar el paquete fileops para escribir y leer en archivos
	"github.com/Pallinder/go-randomdata"
)

const file = "Blanace.txt"

func main() {
	var phone string = randomdata.PhoneNumber()
	fmt.Println("Contacta nuestro centro de servicio", phone)

	var opcion int
	balance, err := fileops.ReadFloatFromFile(file)

	if err != nil {
		fmt.Println(err)
	}

	primerSaludo()

	ingresaOpcion()

	fmt.Scan(&opcion)
	fmt.Println("-----------------------------------------")

	switch opcion {
	case 1:
		bancoIf(balance)
	case 2:
		bancoSwitch(balance)
	}
}

func bancoSwitch(balance float64) {
	for {

		selectorDeOpciones()

		var opcion int

		ingresaOpcion()

		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			fmt.Println(balance)
		case 2:
			fmt.Println("Cuánto dinero quiere depositar?")
		case 3:
			fmt.Println("Cuánto dinero quiere retirar?")
		default:
			fmt.Println("Gracias por preferirnos.")
			return
		}

	}
}

func bancoIf(balance float64) {
	for {

		selectorDeOpciones()

		var opcion int
		fmt.Println("-----------------------------------------")
		fmt.Print("Ingresa la opción: ")
		fmt.Scan(&opcion)

		if opcion == 1 {
			fmt.Println("Tu balance es: ", balance)
			fmt.Println("-----------------------------------------")
		} else if opcion == 2 {
			var valor float64
			fmt.Print("Cuánto vas a depositar: ")
			fmt.Scan(&valor)
			if valor > 0 {
				balance += valor
				fileops.WriteFloatInFile(file, balance)
			} else {
				fmt.Println("No valores negativos.")
				fmt.Println("-----------------------------------------")
				continue
			}
			fmt.Println("-----------------------------------------")
		} else if opcion == 3 {
			var valor float64
			fmt.Print("Cuánto vas a retirar: ")
			fmt.Scan(&valor)
			if balance >= valor && valor > 0 {
				balance -= valor
				fileops.WriteFloatInFile(file, balance)
			} else {
				fmt.Println("No te alcanza.")
				fmt.Println("-----------------------------------------")
				continue
			}
			fmt.Println("-----------------------------------------")
		} else {
			break
		}
	}
	fmt.Println("Gracias por preferirnos.")
}
