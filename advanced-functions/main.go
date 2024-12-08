package main

import (
	"fmt"
	"strconv"
)

// Another way to save the function:
// type transformFunc func(int) int
// func transofrmNumbers(numbers *[]int, trns transformFunc) []int

func main() {

	numbers := []int{1, 2, 3, 4}

	doubleNumbers := func(numbers []int) []int {
		for i, v := range numbers {
			numbers[i] = v + v
		}
		return numbers
	}

	fmt.Println(numbers)
	numbers = doubleNumbers(numbers)
	fmt.Println(numbers)

	numbers = []int{1, 2, 3, 4}

	// If I want to performe a specific operation in the same function,
	// just pass the specific function as argument
	fmt.Println(transofrmNumbers(&numbers, doubleNumber))
	fmt.Println(transofrmNumbers(&numbers, tripleNumber))

	// itoa becomes the function created inside returnFunction()
	itoa := returnFunction("Returning function as value")
	fmt.Println(itoa(8))

}

func transofrmNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func doubleNumber(num int) int {
	return num * 2
}

func tripleNumber(num int) int {
	return num * 3
}

func returnFunction(str string) func(int) string {
	fmt.Println(str)
	myFunc := func(n int) string {
		str := strconv.Itoa(n)
		return str
	}
	return myFunc
}
