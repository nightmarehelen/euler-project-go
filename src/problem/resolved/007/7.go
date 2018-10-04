package main

import "lib/math"

/**
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10 001st prime number?
*/

func main() {

	counter := 0
	for i := int64(2); ; i++{
		if math.IsPrime(i) {
			counter++

			if counter == 10001 {
				print(i)
				return
			}
		}
	}

}
