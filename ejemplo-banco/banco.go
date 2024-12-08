package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "Blanace.txt"

func main() {
	var opcion int
	balance, err := readFromFile()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Hola, bienvenido al banco.")
	fmt.Println("Selecciona una opción:\n1. Banco con if.\n2. Banco con switch.")
	fmt.Scan(&opcion)
	switch opcion {
	case 1:
		bancoIf(balance)
	case 2:
		bancoSwitch(balance)
	}
}

func writeInFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func readFromFile() (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		writeInFile(1000)
		return 1000, errors.New("no existe el archivo")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		writeInFile(1000)
		return 1000, errors.New("el dato no se puede convertir a número")
	}

	return balance, nil
}

func bancoSwitch(balance float64) {
	for {
		fmt.Println("¿Qué quieres hacer?")
		fmt.Println("1. Revisar balance.")
		fmt.Println("2. Despositar dinero.")
		fmt.Println("3. Retirar dinero.")
		fmt.Println("4. Salir")

		var opcion int
		fmt.Println("-----------------------------------------")
		fmt.Print("Ingresa la opción: ")
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

		fmt.Println("¿Qué quieres hacer?")
		fmt.Println("1. Revisar balance.")
		fmt.Println("2. Despositar dinero.")
		fmt.Println("3. Retirar dinero.")
		fmt.Println("4. Salir")

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
				writeInFile(balance)
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
				writeInFile(balance)
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
