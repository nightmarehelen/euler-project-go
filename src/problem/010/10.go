package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"math"
)

func shaishufa() {
	sum := int64(0)
	cha := make(chan int64)
	fd,err := os.OpenFile("primes-less-than-2million.txt", os.O_RDWR| os.O_APPEND, 0644)
	defer fd.Close()
	if err != nil{
		println(err.Error())
		return
	}
	mutex := sync.Mutex{}
	mutex.Lock()
	go generateOriginalSequence(cha)

	//素数公式
	for i := 1; i < 275696; i++ {
		dest := make(chan int64)
		prime := <-cha
		sum += prime
		fmt.Printf("%-9d :%d\n", i, prime)
		fd.Write([]byte(strconv.FormatInt(prime,10)+"\n"))
		if prime > 20000000 {
			mutex.Unlock()
			return
		}

		go filter(cha, dest, prime)
		cha = dest
	}
	mutex.Lock()
	println(sum)

}

func filter(src chan int64, dest chan int64, prime int64) {
	for num := range src {
		if num%prime != 0 {
			dest <- num
		}
	}
}

//generateOriginalSequence 生成原始1-N的数字序列
func generateOriginalSequence(cha chan int64) {

	for i := int64(2); ; i++ {
		cha <- i
	}
}


func IsPrime(num int64) bool {
	if num == 1 {
		return false
	}

	if num == 2 {
		return true
	}

	for i := int64(2); i <= int64(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main(){
	fd,err := os.OpenFile("primes-less-than-2million.txt", os.O_RDWR| os.O_APPEND, 0644)
	defer fd.Close()
	if err != nil {
		println(err.Error())
	}

	var max = int64(2000000)
	sum := int64(2)
	for i:= int64(3); i< max;i = i+2{
		if IsPrime(i){
			fd.Write([]byte(strconv.FormatInt(i,10)+"\n"))
			println(i)
			sum += i
		}
	}

	println(sum)
}