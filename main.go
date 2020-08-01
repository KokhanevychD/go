package main

import (
	"fmt"
)

func main() {
	intArr := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	res1 := Bubblesort(intArr)
	res2 := Insertion(intArr)
	fmt.Println(res1)
	fmt.Println(res2)
}

// Insertion sorting algorithm
func Insertion(list []int) []int {
	for idx := range list {
		if idx > 0 {
			if list[idx] < list[idx-1] {
				list[idx], list[idx-1] = list[idx-1], list[idx]
				// reverse iteration in slice of array, breaks if find place
				slice := list[:idx]
				len := len(slice)
				for idx := range slice {
					if slice[len-idx-1] < slice[len-idx-2] {
						slice[len-idx-1], slice[len-idx-2] = slice[len-idx-2], slice[len-idx-1]
					} else {
						break
					}
				}
			}

		}
	}
	return list
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
