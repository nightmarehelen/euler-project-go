package main
/**
A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4.
If all of the permutations are listed numerically or alphabetically, we call it lexicographic order.
The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
 */

import (
	"fmt"
	"lib/permutation"
)
func main(){

	seq := "0123456789"
	end := false
	for idx := 2; ;idx++{
		seq,end = permutation.Next(seq)
		if end{
			return
		}
		fmt.Println(seq)
		if idx == 1000000{
			fmt.Printf("%s is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9!\n",seq)
			return
		}
	}
}
