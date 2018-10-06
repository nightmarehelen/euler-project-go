package main

import(
	myMathLib "lib/math"
)
/**
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 */
func main(){

	ret := 1
	for i := 1; i<=20; i++{
		ret = myMathLib.Lcm(ret,i)
	}

	print(ret)
}

