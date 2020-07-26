package main

import (
	"fmt"
)

func main(){
	int_arr := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	res := bubblesort(int_arr)
	fmt.Printf("%v", res)
}

func bubblesort(list []int) []int{
	length := len(list)
	for idx := range list{
		if idx != length - 1{
			if list[idx] > list[idx + 1]{
				list[idx], list[idx + 1] = list[idx + 1], list[idx]
				list = bubblesort(list)
			}
		}
	}
	return list
}