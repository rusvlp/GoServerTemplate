package main

import "fmt"

func sandbox() {
	var slice []int = []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	var newSlice []int = []int{5, 5, 6, 4}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
