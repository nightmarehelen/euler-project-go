package main

import (
	"fmt"
	"lib/prime"
)

/**
Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
 */

 func main(){
 	dict := make(map[int64]int64)
 	for i:=int64(2); i<10000; i++{
 		divisors := prime.GetDivisors(i)
 		sum := int64(0)
 		for _,value := range divisors{
 			if value != i{
 				sum = sum + value
			}
		}
		dict[i] = sum
	}

 	sum := int64(0)
 	for key,value := range dict{
 		reverse, ok := dict[value]
 		if ok{
 			if reverse == key && key != value{
 				fmt.Printf("%d\t%d\n",key, value)
				fmt.Printf("%d=%v\n",key, prime.GetDivisors(key))
				fmt.Printf("%d=%v\n",value, prime.GetDivisors(value))
 				sum = sum + key
			}
		}
	}
 	fmt.Println(sum)
 }
