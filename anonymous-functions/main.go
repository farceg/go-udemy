package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	fmt.Println(double(8))

	triple := createTransformer(3)
	fmt.Println(triple(8))

	// When the function is needed as a value, one can create it there
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformed)

	transformed = transformNumbers(&numbers, func(number int) int {
		return number * 3
	})

	fmt.Println(transformed)

	transformed = transformNumbers(&numbers, double)

	fmt.Println(transformed)

	transformed = transformNumbers(&numbers, triple)

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(num int) int {
		return num * factor
	}
}
