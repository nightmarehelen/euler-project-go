package permutation

import (
	"lib/myorder"
	"math"
)
//Next 返回下一个序列，end表示是否为随后一个序列
func Next(input string) (ret string, end bool){
	idx := lastInvertedSequence(input)
	if idx == -1{
		return input,true
	}
	//找到后面序列大于当前序列的最小的位
	idx2 := getMinLargerThan(input[idx], input[idx+1:])
	bytes := []byte(input)
	//交换次序
	temp := bytes[idx]
	bytes[idx] = bytes[idx+1:][idx2]
	bytes[idx+1:][idx2] = temp
	head := bytes[:idx+1]
	so := myorder.StringOrder{Value:string(bytes[idx+1:])}
	return string(head) + (&so).Sort().Value, false
}

//lastInvertedSequence 返回末尾的逆序的前一个索引
func lastInvertedSequence(input string) int {
	for i := len(input) - 1; i >= 1; i--{
		if input[i] > input[i-1] {
			return i-1
		}
	}
	//已经按照从大到小的顺序排列
	return -1
}

//getMinLargerThan 获取序列当中大于制定元素的最小的元素
func getMinLargerThan(value uint8, input string) int{
	min := uint8(math.MaxUint8)
	index := -1
	for idx, tempValue := range input{
		if uint8(tempValue) > value && uint8(tempValue) < min{
			min = uint8(tempValue)
			index = idx
		}
	}
	return index
}
