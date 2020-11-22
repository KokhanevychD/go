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
	res4 := problem.ByState(
		"John Daggett, 341 King Road, Plymouth MA \nAlice Ford, 22 East Broadway, Richmond VA \nOrville Thomas, 11345 Oak Bridge Road, Tulsa OK \nTerry Kalkas, 402 Lans Road, Beaver Falls PA \nEric Adams, 20 Post Road, Sudbury MA \nHubert Sims, 328A Brook Road, Roanoke VA \nAmy Wilde, 334 Bayshore Pkwy, Mountain View CA \nSal Carpenter, 73 6th Street, Boston MA \n")
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	problem.PlusMinus(int32arr)
	fmt.Println(res4)
}
