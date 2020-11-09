package main

import (
	"fmt"

	"github/go/algo"
	"github/go/problem"
)

func main() {
	intArr := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	int32arr := []int32{-4, 3, -9, 0, 4, 1}
	res1 := algo.Bubblesort(intArr)
	res2 := algo.Insertion(intArr)
	res3 := algo.ReverseVowels("Hello World")
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	problem.PlusMinus(int32arr)
}
