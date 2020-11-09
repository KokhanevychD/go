package problem

import "fmt"

// PlusMinus - Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero. Print the decimal value of each fraction on a new line with  places after the decimal. HackerRank problem.
func PlusMinus(arr []int32) {
	positive := 0
	negative := 0
	zero := 0
	for _, num := range arr {
		if num > 0 {
			positive++
		} else if num < 0 {
			negative++
		} else {
			zero++
		}
	}
	list := []int{positive, negative, zero}
	for _, num := range list {
		res := fmt.Sprintf("%.6f", float64(num)/float64(len(arr)))
		fmt.Println(res)
	}
}
