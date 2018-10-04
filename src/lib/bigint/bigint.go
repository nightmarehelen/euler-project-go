package bigint

import (
	"strings"
)

type BigInt struct{
	len int
	digits []byte
}

func (b *BigInt) Digits() []byte{
	return b.digits
}

func (b *BigInt) Len() int{
	return b.len
}

//New 构造函数
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

//Compare 比较大小
func (b *BigInt) Compare(bigInt *BigInt) int{
	if b.len > bigInt.len{
		return 1
	}else if b.len < bigInt.len{
		return -1
	}else{
		for i :=0; i<b.len; i++{
			if b.digits[i] < bigInt.digits[i]{
				return -1
			}else if b.digits[i] > bigInt.digits[i]{
				return 1
			}
		}
		return 0
	}
}

//Subtract 减去参数，要求被减数大于减数
func (b *BigInt) Subtract(bigInt *BigInt) *BigInt{
	if b.Compare(bigInt) == -1{
		panic("要求被减数大于减数")
	}
	subtractor := make([]byte, b.len, b.len)
	//对齐
	for i:=0; i<bigInt.len; i++{
		subtractor[b.len - bigInt.len + i] = bigInt.digits[i]
	}
	borrow := byte(0)
	ret := make([]byte, b.len, b.len)
	for i:=b.len-1; i>=0; i--{
		if b.digits[i] >= subtractor[i] + borrow {
			ret[i] = b.digits[i] - subtractor[i] - borrow
			borrow = 0
		}else{
			ret[i] = 10 + b.digits[i] - subtractor[i] - borrow
			borrow = 1
		}
	}
	for i:=0; i<b.len; i++{
		if ret[i] != 0{
			return &BigInt{len(ret[i:]), ret[i:]}
		}
	}
	return &BigInt{len(ret), ret}
}


//MultiplySingle 大整数乘以单个数字
func (b *BigInt) MultiplySingle(digit byte) *BigInt{
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

//Multiply 乘法
func (b *BigInt) Multiply(bigInt *BigInt) *BigInt{
	result :=New("0")
	for i:=0; i<bigInt.len; i++{
		temp := b.MultiplySingle(bigInt.digits[bigInt.len-1-i])
		for j:=0; j<i; j++{
			temp.digits = append(temp.digits, 0)
			temp.len = temp.len + 1
		}
		result = result.Add(temp)
	}
	return result
}

//Divided 除法，除以参数,要求被除数大于除数
func (b * BigInt)Divide(bigInt *BigInt) (result *BigInt, remainder *BigInt){

	//结果最多b.len - bigInt.len位
	ret := make([]byte,	b.len - bigInt.len + 1, b.len - bigInt.len +1)
	result = &BigInt{b.len - bigInt.len + 1, ret}

	temp := &BigInt{b.len, b.digits}
	for i:=b.len - bigInt.len; i>=0;i--{
		pow :=New("1")
		for j:=0; j<i;j++{
			pow = pow.Multiply(New("10"))
		}

		for j:=byte(0); j<=9;j++{
			if pow.MultiplySingle(j+1).Multiply(bigInt).Compare(temp) > 0{
				result.digits[result.len-i-1] = j
				temp = temp.Subtract(pow.MultiplySingle(j).Multiply(bigInt))
				break
			}else{
				result.digits[result.len-i -1] = 0
			}
		}
	}

	return result, b.Subtract(result.Multiply(bigInt))
}