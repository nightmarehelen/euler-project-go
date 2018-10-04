package main

import "lib/prime"

func main(){
	primes := prime.New(1000)
	primes.Init()
	primes.Dump()
}
