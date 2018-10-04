package main

import (
	"fmt"
	"os"
)

func main1() {
	sum := int64(0)
	for i := int64(1); ; i++ {
		sum = sum + i
		getDivisorCount(sum)
	}
}

func getDivisorCount(num int64){
	cnt := int64(0)
	for i := int64(1); i <= num; i++ {
		if num%i == 0 {
			cnt = cnt + 1
		}

	}

	fmt.Printf("num = %d, cnt= %d\n", num, cnt)
	if cnt >= 500 {

		fmt.Printf("final result num = %d, cnt= %d\n", num, cnt)
		os.Exit(-1)
		return
	}
}
