package bigint

import (
	"strings"
)

type BigInt struct{
	len int
	digits []byte
}


func New(number string) *BigInt{

	digits := make([]byte, 0,0)
	for _,digit := range number{
		digits = append(digits, byte(digit - '0'))
	}
	return &BigInt{len(digits),digits}
}

func (b * BigInt) String() string{
	output := strings.Builder{}
	for _,value := range b.digits{
		output.WriteByte(value+'0')
	}
	return output.String()
}

//add 两个正整数的加法
func (b *BigInt) Add(num *BigInt) *BigInt{
	//构造新的整数，两端对齐
	var bBytes []byte = nil
	var numBytes []byte = nil
	if b.len >= num.len{
		bBytes = make([]byte, b.len+1, b.len+1)
		bBytes[0] = 0
		for idx,value := range b.digits{
			bBytes[idx+1] = value
		}
		numBytes = make([]byte, b.len+1, b.len+1)
		numBytes[0] = 0
		for idx,value := range num.digits{
			numBytes[b.len - num.len + idx+1] = value
		}
	}else{
		numBytes = make([]byte, num.len+1, num.len+1)
		numBytes[0] = 0
		for idx,value := range num.digits{
			numBytes[idx+1] = value
		}
		bBytes = make([]byte, num.len+1, num.len+1)
		bBytes[0] = 0
		for idx,value := range b.digits{
			bBytes[num.len - b.len + idx + 1] = value
		}
	}

	var ret = make([]byte, len(bBytes), len(bBytes))
	carry := byte(0)
	for i := len(bBytes)-1; i>=0; i--{
		temp := bBytes[i] + numBytes[i] + carry
		if temp >= 10{
			ret[i] = temp %10
			carry = 1
		}else{
			ret[i] = temp
			carry = 0
		}
	}

	if ret[0] == 0{
		ret = ret[1:]
	}
	return &BigInt{len:len(ret), digits:ret}
}

//multiplySingle 大整数乘以单个数字
func (b *BigInt) multiplySingle(digit byte) *BigInt{
	digits := make([]byte, b.len+1, b.len+1)
	digits[0] = 0

	carry := byte(0)
	for i:= b.len-1; i>=0; i--{
		temp := b.digits[i] * digit + carry

		if temp >= 10{
			digits[1+i] = temp%10
			carry = byte(temp/10)
		}else{
			digits[1+i] = temp
			carry = 0
		}
	}
	digits[0] = carry
	if digits[0] != 0 {
		return &BigInt{len(digits), digits}
	}else{
		return &BigInt{len(digits) -1 , digits[1:]}
	}
}

//
func (b *BigInt) Multiply(bigInt *BigInt) *BigInt{
	result :=New("0")
	for i:=0; i<bigInt.len; i++{
		temp := b.multiplySingle(bigInt.digits[bigInt.len-1-i])
		for j:=0; j<i; j++{
			temp.digits = append(temp.digits, 0)
			temp.len = temp.len + 1
		}
		result = result.Add(temp)
	}
	return result
}
