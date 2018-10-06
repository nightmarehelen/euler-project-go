package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
A palindromic number reads the same both ways.
The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
 */
func main(){

	ret :="0"
	for i :=100; i<=999; i++{
		for j :=100;j<=999; j++{
			temp := strconv.Itoa(i *j)

			isPalindromic := true
			length := len(temp)
			for k :=0 ; k <= length/2; k++{
				if temp[k] != temp[length-1-k]{
					isPalindromic = false
					break
				}
			}
			if isPalindromic {
				fmt.Printf("j=%d\tj=%d\ttemp=%s\n",i,j,temp)

				if len(ret) < len(temp) || strings.Compare(ret,temp) <0{
					ret =temp
				}
			}
		}
	}
	println("\n\n\n\n"+ret)
}
