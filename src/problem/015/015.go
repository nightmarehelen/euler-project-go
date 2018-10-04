package main

import (
	"fmt"
	"lib/bigint"
	"strconv"
)

func main(){
	a := continuedProduct(40)
	b := continuedProduct(20)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a.Divide(b.Multiply(b)))
}

func continuedProduct(num int) *bigint.BigInt{
	ret := bigint.New("1")
	for i:=2;i<=num;i++{
		ret = ret.Multiply(bigint.New(strconv.Itoa(i)))
	}
	return ret
}
