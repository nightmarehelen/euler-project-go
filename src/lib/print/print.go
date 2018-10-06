package print

import (
	"fmt"
	"math"
	"strconv"
)

func TrianglePrint(str []string, space int){
	layer := int((-1+int64(math.Sqrt(float64(1+4*1*2*len(str)))))/2)
	if layer * (layer+1)/2 != len(str){
		panic("the input array is not a triangle sequence")
	}

	//最大长度
	lineLength := layer + layer -1
	for i:=1; i<=layer; i++{
		//行首空格数
		spaces := (lineLength - i - (i-1))/2
		for j:=0; j<spaces; j++{
			fmt.Printf(fmt.Sprintf("%%%ds", space), "")
		}

		printed := (i-1)*i/2
		for k:=0; k<i; k++{
			//fmt.Printf("i=%d\tprinted=%d\tk=%d",i,printed,k)
			fmt.Printf(fmt.Sprintf("%%%ds", space), str[printed+k])
			fmt.Printf(fmt.Sprintf("%%%ds", space), "")
		}

		fmt.Println()
	}
}

type ValueAndPath struct{
	Value int
	Path []int
	OriginValue []int
}

func TrianglePathPrint(layer int, space int, path []int, originValues []int){
	//最大长度
	lineLength := layer + layer -1

	strs := make([][]string, layer, lineLength)
	for i:=0; i<layer; i++{
		strs[i] = make([]string, lineLength, lineLength)
		for j:=0; j<lineLength; j++{
			strs[i][j] = fmt.Sprintf(fmt.Sprintf("%%%ds", space), "--")
		}
	}


	for i:=1; i<=len(path); i++{
		idx := path[i-1]
		if idx >0{
			strs[i-1][(lineLength-i-(i-1))/2+(idx-1)*2+1-1] = strconv.Itoa(originValues[i-1])
		}
	}

	for i:=0; i<len(path); i++{
		for j:=0; j<lineLength; j++{
			fmt.Printf(fmt.Sprintf("%%%ds", space), strs[i][j])
		}
		fmt.Println()
	}
}

func TriangleValuePrint(str []ValueAndPath, space int){
	layer := int((-1+int64(math.Sqrt(float64(1+4*1*2*len(str)))))/2)
	if layer * (layer+1)/2 != len(str){
		panic("the input array is not a triangle sequence")
	}

	//最大长度
	lineLength := layer + layer -1
	for i:=1; i<=layer; i++{
		//行首空格数
		spaces := (lineLength - i - (i-1))/2
		for j:=0; j<spaces; j++{
			fmt.Printf(fmt.Sprintf("%%%ds", space), "")
		}

		printed := (i-1)*i/2
		for k:=0; k<i; k++{
			fmt.Printf(fmt.Sprintf("%%%dd", space), str[printed+k].Value)
			fmt.Printf(fmt.Sprintf("%%%ds", space), "")
		}

		fmt.Println()
	}
}