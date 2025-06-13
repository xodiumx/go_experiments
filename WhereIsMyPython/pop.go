package main

import "fmt"

func pop(slice []int) (int, []int) {
	if len(slice) == 0 {
		return 0, slice
	}
	lastElement := slice[len(slice)-1]
	newSlice := slice[:len(slice)-1]
	return lastElement, newSlice
}

func main() {
	myList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var elem int

	for len(myList) > 0 {
		elem, myList = pop(myList)
		fmt.Println(elem, myList)
	}
}
