package main

import (
	"fmt"
	"math/big"
)

func main() {
	//fibonacciCall()
	factorialCall()
}

func fibonacciCall() {
	n := 45

	// Create channels to receive results from fib and fibDynamicProgramming
	fibChan := make(chan int)
	fibDynamicChan := make(chan int)

	// Start goroutines for fib and fibDynamicProgramming
	go func() {
		result := fib(n)
		fibChan <- result
	}()
	go func() {
		result := fibDynamicProgramming(n)
		fibDynamicChan <- result
	}()

	// Wait for the first result to arrive and print it
	select {
	case result := <-fibChan:
		fmt.Println("Recursive:", result)
	case result := <-fibDynamicChan:
		fmt.Println("Dynamic Programming:", result)
	}

	// Wait for the second result to arrive and print it
	select {
	case result := <-fibChan:
		fmt.Println("Recursive:", result)
	case result := <-fibDynamicChan:
		fmt.Println("Dynamic Programming:", result)
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fibDynamicProgramming(n int) int {
	if n < 2 {
		return n
	}
	fibArray := make([]int, n+1)
	fibArray[0] = 0
	fibArray[1] = 1
	for i := 2; i <= n; i++ {
		fibArray[i] = fibArray[i-1] + fibArray[i-2]
	}
	return fibArray[n]
}

func factorialCall() {
	var n int64 = 55

	// Create channels to receive results from fib and fibDynamicProgramming
	factorialChan := make(chan *big.Int)
	factorialDynamicChan := make(chan *big.Int)

	// Start goroutines for fib and fibDynamicProgramming
	go func() {
		var result *big.Int = factorial(n)
		factorialChan <- result
	}()
	go func() {
		var result *big.Int = factorialDynamicProgramming(n)
		factorialDynamicChan <- result
	}()

	// Wait for the first result to arrive and print it
	select {
	case result := <-factorialChan:
		fmt.Println("Recursive:", result)
	case result := <-factorialDynamicChan:
		fmt.Println("Dynamic Programming:", result)
	}

	// Wait for the second result to arrive and print it
	select {
	case result := <-factorialChan:
		fmt.Println("Recursive:", result)
	case result := <-factorialDynamicChan:
		fmt.Println("Dynamic Programming:", result)
	}
}

func factorial(n int64) *big.Int {
	if n < 0 {
		panic("Factorial is not defined for negative numbers.")
	}

	if n == 0 {
		return big.NewInt(1)
	}

	result := new(big.Int).SetInt64(n)
	return result.Mul(result, factorial(n-1))
}

func factorialDynamicProgramming(n int64) *big.Int {
	if n < 0 {
		panic("Factorial is not defined for negative numbers.")
	}

	dp := make([]*big.Int, n+1)

	dp[0] = big.NewInt(1)

	for i := int64(1); i <= n; i++ {
		dp[i] = new(big.Int).Mul(dp[i-1], big.NewInt(i))
	}

	return dp[n]
}
