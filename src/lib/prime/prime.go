package prime

import (
	"fmt"
	"math"
	"strings"
)

//Prime 素数类
type Prime struct{
	//总个数
	Cnt int
	//最大素数
	maxPrime int64
	//素数数组
	primes []int64
}

type Factor struct{
	prime int64
	power int
}
//builder
func New(cnt int) *Prime{
	return &Prime{Cnt: cnt}
}

//初始化素数数组
func (p *Prime) Init() *Prime{
	num := int64(2)
	for ;len(p.primes) < p.Cnt;{
		if IsPrime(num){
			p.primes = append(p.primes,num)
			p.maxPrime = num
		}
		num = num + 1
	}
	return p
}

//输出素数信息
func (p *Prime) Dump(){
	fmt.Println("Dump prime Info")
	fmt.Printf("Prime Count %d\n",p.Cnt)
	fmt.Printf("Max Prime %d\n", p.maxPrime)
	fmt.Println("Primes List:")
	fmt.Println(p.primes)
}

//IsPrime 判断一个数是否为素数
func IsPrime(num int64) bool{
	for i := int64(2); i<= int64(math.Sqrt(float64(num))); i++{
		if num % i == 0{
			return false
		}
	}
	return true
}

//返回最大的素数
func (p *Prime) MaxPrime() int64{
	return p.maxPrime
}

//返回素数数组
func (p *Prime) Primes() []int64{
	return p.primes
}

//返回指定范围内的素数
func (p *Prime) PrimeArrayByStartEnd(start, end int64) []int64{
	idxStart,idxEnd := -1,-1
	for i := 0; i<len(p.primes); i++{
		if start == p.primes[i] && idxStart == -1{
			idxStart = i
		}

		if end == p.primes[i] && idxEnd == -1{
			idxEnd = i
		}
	}
	return p.primes[idxStart : idxEnd + 1]
}

//返回从2开始的制定长度的素数
func (p *Prime) PrimeArrayByLength(length int) []int64{
	return p.primes[0 : length]
}


//返回小于指定数据的所有素数
func (p *Prime)PrimeArrayBelow(number int64){
	if number > p.maxPrime{
		for i:=p.maxPrime+1; i<= number; i++{
			if IsPrime(i){
				p.primes = append(p.primes, i)
				p.maxPrime = i
				p.Cnt = len(p.primes)
			}
		}
	}
}

//isQuickPrime 快速判断一个数是否为素数，利用素数表
func (p *Prime) isQuickPrime(number int64) bool{
	if number < p.maxPrime{
		for _,value := range p.primes{

			if number == value{
				return true
			}else if number % value == 0{
				return false
			}
		}
		return true
	}

	if number == p.maxPrime{
		return true
	}

	for i := p.maxPrime + 1; i<= int64(math.Sqrt(float64(number))); i++{
		if number % i == 0{
			return false
		}
	}

	return true
}

// 分解质因数
func (p *Prime) DecompositionPrimeFactor(number int64) []Factor{

	p.PrimeArrayBelow(number)


	factors  := make([]Factor, 0, 0)
	//1.首先找到素数表当中的质因子的幂次
	for _,prime := range p.primes{
		//有该质因子
		if number % prime == 0{
			factor := decompositionPrimeFactor(number, prime)
			factors = append(factors, *factor)
		}
	}

	//格式化输出质因子
	output := strings.Builder{}
	for _,value := range factors{
		output.WriteString(fmt.Sprintf("%d^^%d*", value.prime, value.power))
	}
	fmt.Printf(" %d = %s ",number, output.String())
	return factors
}



//单个质因子幂次计算
func decompositionPrimeFactor(number int64, prime int64) *Factor {
	for i := 1; ; i++{
		if number % int64(math.Pow(float64(prime), float64(i+1))) != 0{
			return &Factor{prime, i}
		}
	}
}

//获取约数个数
func (p *Prime) GetDivisorCount(number int64) int{
	factors := p.DecompositionPrimeFactor(number)

	cnt := 1
	for _,value := range factors{
		cnt  = cnt * (value.power + 1)
	}
	return cnt
}



