package main

import "fmt"

/**
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
 */

 func main(){
 	for i:=1;i<1000;i++{
 		for j:=1;j<1000;j++{
 			for k:=1;k<1000;k++{
 				if i+j+k == 1000 && i*i + j*j == k*k{
 					fmt.Printf("i=%4d,j=%4d,k=%4d\n", i, j, k)
 					println(i * j * k)
				}
			}
		}
	}
 }
