package main

import (
	"fmt"
	"lib/bigint"
	"strconv"
)

/**
215 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 21000?
 */
func main(){
	product := powerOf2(15)
	fmt.Println(product.String())
	product = powerOf2(1000)
	fmt.Println(product.String())
	sum := 0
	for i:=0;i< product.Len();i++{
		sum = sum + int(product.Digits()[i])
	}
	fmt.Println(sum)
}

func powerOf2(num int) *bigint.BigInt{
	ret := bigint.New("1")

	for i:=0; i<num; i++{
		ret = ret.Multiply(bigint.New(strconv.Itoa(2)))
	}
	return ret
}
