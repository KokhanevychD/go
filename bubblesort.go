package main

import (
	"fmt"
)

func main(){
	int_arr := []int{123, 1425, 3462, 7, 347, 6432, 312}
	res := bublesort(int_arr)
	fmt.Printf("%v", res)
}

func bublesort(list []int) []int{
	length := len(list)
	for idx := range list{
		if idx != length - 1{
			if list[idx] > list[idx + 1]{
				list[idx], list[idx + 1] = list[idx + 1], list[idx]
				list = bublesort(list)
			}
		}
	}
	return list
}