package main

import (
	"fmt"
	"io/ioutil"
	"lib/print"
	"math"
	"strconv"
	"strings"
)

/**
By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 023.

3
7 4
2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 023.

Find the maximum total from top to bottom of the triangle below:

75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 023 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 023 09 70 98 73 93 38 53 60 04 023

NOTE: As there are only 16384 routes, it is possible to solve this problem by trying every route.
However, Problem 67, is the same challenge with a triangle containing one-hundred rows;
it cannot be solved by brute force, and requires a clever method! ;o)
 */
func main() {
	content, err := ioutil.ReadFile("src/file/18.txt")
	if err != nil {
		panic(err)
	}

	noLineSep := strings.Replace(string(content), "\n", " ", -1)
	numbers := strings.Split(noLineSep, " ")
	//fmt.Println(numbers)
	//fmt.Println(len(numbers))
	//fmt.Println(15*(1+15)/2)
	print.TrianglePrint(numbers, 2)

	valueAndPaths := make([]print.ValueAndPath, len(numbers), len(numbers))

	layer := int((-1 + int64(math.Sqrt(float64(1+4*1*2*len(valueAndPaths))))) / 2)
	if layer*(layer+1)/2 != len(valueAndPaths) {
		panic("the input array is not a triangle sequence")
	}

	//初始化节点值
	for idx, value := range numbers {
		num, err := strconv.Atoi(value)
		if err != nil {
			panic("string convert to number error!")
		}
		valueAndPaths[idx] = print.ValueAndPath{Value: num}
		valueAndPaths[idx].Path = make([]int, layer, layer)
		valueAndPaths[idx].OriginValue = make([]int, layer, layer)
	}

	MaxTotal(numbers, valueAndPaths)
}

func MaxTotal(numbers []string, valueAndPaths []print.ValueAndPath) {
	layer := int((-1 + int64(math.Sqrt(float64(1+4*1*2*len(valueAndPaths))))) / 2)
	if layer*(layer+1)/2 != len(valueAndPaths) {
		panic("the input array is not a triangle sequence")
	}

	valueAndPaths[0].Path[0] = 1
	valueAndPaths[0].OriginValue[0] = valueAndPaths[0].Value
	//每层
	for i := 2; i <= layer; i++ {
		for j := 1; j <= i; j++ {
			leftTop := (i-2)*(i-2+1)/2 + j - 1 - 1
			rightTop := (i-2)*(i-2+1)/2 + j - 1
			current := (i-1)*(i-1+1)/2 + j - 1


			fmt.Printf("i=%d\tj=%d\tlefttop=%d\trighttop=%d\tcurrent=%d\n", i, j, leftTop, rightTop, current)
			if leftTop>=0 && leftTop <len(valueAndPaths){
				fmt.Printf("lefttop=%d\n",  valueAndPaths[leftTop])
			}

			if rightTop>=0 && rightTop <len(valueAndPaths){
				fmt.Printf("rightTop=%d\n",  valueAndPaths[rightTop])
			}


			if j == 1 {
				if i != 1 {
					fmt.Println("rightTop")
					print.TrianglePathPrint(layer, 2, valueAndPaths[rightTop].Path, valueAndPaths[rightTop].OriginValue)
					fmt.Println(valueAndPaths[rightTop].Path)

					copy(valueAndPaths[current].OriginValue[0:i-1], valueAndPaths[rightTop].OriginValue[0:i-1])
					valueAndPaths[current].OriginValue[i-1] = valueAndPaths[current].Value
					copy(valueAndPaths[current].Path[0:i-1], valueAndPaths[rightTop].Path[0:i-1])
					valueAndPaths[current].Path[i-1] = j

					valueAndPaths[current].Value = valueAndPaths[rightTop].Value + valueAndPaths[current].Value

				}
			} else if j == i {
				fmt.Println("leftTop")
				print.TrianglePathPrint(layer, 2, valueAndPaths[leftTop].Path, valueAndPaths[leftTop].OriginValue)
				fmt.Println(valueAndPaths[leftTop].Path)


				copy(valueAndPaths[current].OriginValue[0:i-1], valueAndPaths[leftTop].OriginValue[0:i-1])
				valueAndPaths[current].OriginValue[i-1] = valueAndPaths[current].Value
				copy(valueAndPaths[current].Path[0:i-1], valueAndPaths[leftTop].Path[0:i-1])
				valueAndPaths[current].Path[i-1] = j

				valueAndPaths[current].Value = valueAndPaths[leftTop].Value + valueAndPaths[current].Value
			} else {
				if valueAndPaths[leftTop].Value >= valueAndPaths[rightTop].Value {
					fmt.Println("leftTop")
					print.TrianglePathPrint(layer, 2, valueAndPaths[leftTop].Path, valueAndPaths[leftTop].OriginValue)
					fmt.Println(valueAndPaths[leftTop].Path)

					copy(valueAndPaths[current].OriginValue[0:i-1], valueAndPaths[leftTop].OriginValue[0:i-1])
					valueAndPaths[current].OriginValue[i-1] = valueAndPaths[current].Value
					copy(valueAndPaths[current].Path[0:i-1], valueAndPaths[leftTop].Path[0:i-1])
					valueAndPaths[current].Path[i-1] = j

					valueAndPaths[current].Value = valueAndPaths[leftTop].Value + valueAndPaths[current].Value
				} else {
					fmt.Println("rightTop")
					print.TrianglePathPrint(layer, 2, valueAndPaths[rightTop].Path, valueAndPaths[rightTop].OriginValue)
					fmt.Println(valueAndPaths[rightTop].Path)

					copy(valueAndPaths[current].OriginValue[0:i-1], valueAndPaths[rightTop].OriginValue[0:i-1])
					valueAndPaths[current].OriginValue[i-1] = valueAndPaths[current].Value
					copy(valueAndPaths[current].Path[0:i-1], valueAndPaths[rightTop].Path[0:i-1])
					valueAndPaths[current].Path[i-1] = j
					valueAndPaths[current].Value = valueAndPaths[rightTop].Value + valueAndPaths[current].Value
				}
			}

			fmt.Println("valueAndPaths")
			print.TriangleValuePrint(valueAndPaths, 4)
			fmt.Println("TrianglePrint")
			print.TrianglePrint(numbers, 2)

			fmt.Println(valueAndPaths[current].Path)
			fmt.Println("TrianglePathPrint")
			print.TrianglePathPrint(layer, 2, valueAndPaths[current].Path, valueAndPaths[current].OriginValue)
			fmt.Println("******************************************************************************************************************************")
		}
	}

	maxIdx, maxValue := -1, -1
	for i := 1; i <= layer; i++ {
		if valueAndPaths[ (layer-1)*(layer-1+1)/2+i-1].Value >= maxValue {
			maxIdx = (layer-1)*(layer-1+1)/2 + i - 1
			maxValue = valueAndPaths[ (layer-1)*(layer-1+1)/2+i-1].Value
		}
	}
	fmt.Printf("maxValue = %d", maxValue)
	fmt.Println("Path is as follow:")
	fmt.Println(valueAndPaths[maxIdx].OriginValue)
	fmt.Println(valueAndPaths[maxIdx].Path)
	print.TrianglePathPrint(layer, 2, valueAndPaths[maxIdx].Path, valueAndPaths[maxIdx].OriginValue)
}
