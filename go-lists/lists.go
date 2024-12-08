package main

import "fmt"

type product struct {
	id    int32
	title string
	price float64
}

func main() {

	// Array
	var newArr [4]int
	fmt.Println(newArr, "(declared array, and automatically intialized to zero values)")

	newArr[3] = 4
	fmt.Println(newArr, "(modified the index 3 of the array)")

	// Slice (window to the array)
	sliceOfArr := newArr[:2]
	sliceOfArr[0] = 1
	fmt.Println(sliceOfArr, "(a slice is a portion of an array)", len(sliceOfArr), cap(sliceOfArr))
	fmt.Println(newArr, "(the array also will be modified if the slice is)", len(newArr), cap(newArr))

	// Dynamic array
	prices := []float64{5.89, 4.41}
	fmt.Println(prices, "(fixed size of the array with some values)")
	prices = append(prices, 1.23)
	fmt.Println(prices, "(the array grows using the append function)")

	// List o products
	products := []product{}
	products = append(products, product{1, "Pencil", 5.8})
	products = append(products, product{2, "Eraser", 4.1})
	fmt.Println(products)

	discountProducts := []product{{3, "Marker", 2.1}, {4, "Sharpener", 1.4}}

	products = append(products, discountProducts...)
	fmt.Println(products)

}
