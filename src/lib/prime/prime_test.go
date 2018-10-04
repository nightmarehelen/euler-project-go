package prime

import (
	"fmt"
	"testing"
)


func TestInit(t *testing.T){
	prime := New(100)
	prime.Init()
	prime.Dump()

	fmt.Println(prime.MaxPrime())
	fmt.Println(prime.Primes())
	fmt.Println(prime.PrimeArrayByStartEnd(2,61))
	fmt.Println(prime.PrimeArrayByLength(10))
}

func TestDecompositPrimeFactor(t *testing.T){
	prime := New(0)

	prime.DecompositionPrimeFactor(120)

	if 16 != prime.GetDivisorCount(120){
		t.Log("getDivisorCount fail")
	}

}
