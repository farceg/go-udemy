package main

import "fmt"

type example struct {
	name string
	age  int
}

func main() {
	newExample := example{name: "Felipe", age: 28}
	printAnything("Hola")
	printAnything(548)
	printAnything(true)
	printAnything(newExample)
	printAnything(45.9)
	fmt.Println("-------------------------------------")
	printAnything(add(5, 9))
	printAnything(add(0.1, 9.9))
	printAnything(add("Hola", "Mundo"))
}

// También se puede llamar 'any' en reemplazo de 'interface{}'
func printAnything(value interface{}) any {
	switch value.(type) {
	case string:
		fmt.Printf("String value: %s\n", value)
	case int:
		fmt.Printf("Int value: %d\n", value)
	case bool:
		fmt.Printf("Int value: %t\n", value)
	case example:
		fmt.Printf("Struct example: %+v\n", value)
	default:
		valor := fmt.Sprintf("%T", value)
		fmt.Printf("Another value: %v of type: %s\n", value, valor)
	}
	return value
}

// T significa que se puede usar esta función para cualquier tipo de dato
// Pero en este caso se fa una lista de datos posibles, pues no todos permiten el símbolo '+'
func add[T int | float64 | string](a, b T) T {
	return a + b
}
