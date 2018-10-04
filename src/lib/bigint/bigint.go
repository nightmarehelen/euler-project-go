package bigint

import "strings"

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
