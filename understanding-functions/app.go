package main

import "fmt"

func main() {
	fmt.Println("Defer:")
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		defer fmt.Println(i)
		fmt.Println(i)
	}
}
