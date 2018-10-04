package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"lib/math"
	"time"
)

/**
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
 */
func main(){
	t1 := time.Now()
	for i := int64(2) ;i<=int64(600851475143); i++{
		if 600851475143%i ==0 && math.IsPrime(600851475143/i){
			print(600851475143/i)
			elapsed := time.Since(t1)
			fmt.Println("App elapsed: ", elapsed)
			return
		}
	}
}

func findAllPrimes(){
	maxInt64,err := strconv.ParseInt("7FFFFFFFFFFFFFFF",16,64)

	if err != nil{
		panic("can not found max int64")
	}

	println(maxInt64)

	//maxInt64 = 1000

	filename := "primes.txt"
	var primes []int64
	for i := int64(2) ; i < maxInt64;i++{
		if math.IsPrime(i) {
			println(i)
			primes = append(primes, i)
		}
	}
	retStr :=""
	for i:=0; i<len(primes);i++{
		retStr += strconv.FormatInt(primes[i],10)+"\n"
	}
	ioutil.WriteFile(filename, []byte(retStr) , 777)
}
