package main

import "fmt"

/**
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
 */

func getSequenceCount(registed map[int64]int ,num int64) int{
	cnt := 0
	for ;num != 1;{
		old,ok := registed[num]
		if ok {
			cnt = old + cnt
			break
		}
		if num % 2 ==0{
			num = num/2

		}else{
			num = 3*num +1
		}
		cnt ++
	}
	registed[num] = cnt
	return cnt
}

func main(){
	registered := make(map[int64]int)
	maxCnt := 0
	maxNum := int64(0)
	for i := int64(2); i<= 1000000; i++{
		cnt := getSequenceCount(registered , i)

		if cnt > maxCnt{
			maxCnt = cnt
			maxNum = i
		}
	}
	fmt.Printf("num = %d,cnt = %d\n", maxNum, maxCnt)
}
