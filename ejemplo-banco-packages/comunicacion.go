package main

import "fmt"

func primerSaludo() {
	fmt.Println("Hola, bienvenido al banco.")
	fmt.Println("Selecciona una opción:\n1. Banco con if.\n2. Banco con switch.")
}

func selectorDeOpciones() {
	fmt.Println("¿Qué quieres hacer?")
	fmt.Println("1. Revisar balance.")
	fmt.Println("2. Despositar dinero.")
	fmt.Println("3. Retirar dinero.")
	fmt.Println("4. Salir")
}

func ingresaOpcion() {
	fmt.Println("-----------------------------------------")
	fmt.Print("Ingresa la opción: ")
}
