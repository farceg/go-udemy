package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := append(a[:1], 10, 11, 12)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	fmt.Println("--------------------")

	c := []int{1, 2, 3, 4}
	d := append(c[:1], 10, 11)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
}
