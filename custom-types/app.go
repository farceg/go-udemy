package main

import "fmt"

type customString string

func (s customString) log() {
	fmt.Println(s)
}

func main() {
	var name customString
	name = "Felipe"
	name.log()

	strValue := customString("Juan")
	strValue.log()
}
