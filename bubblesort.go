package main

import (
	"fmt"
)

func main() {
	intArr := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	res := Bubblesort(intArr)
	fmt.Printf("%v", res)
}

// Bubblesort fun, sorting int arrays
func Bubblesort(list []int) []int {
	length := len(list)
	sorted := true
	for idx := range list {
		if idx != length-1 {
			if list[idx] > list[idx+1] {
				list[idx], list[idx+1] = list[idx+1], list[idx]
				sorted = false
			}
		}
	}
	if !sorted {
		list = Bubblesort(list)
	}
	return list
}
