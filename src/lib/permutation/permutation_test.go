package permutation

import (
	"fmt"
	"testing"
)

func TestNext(t *testing.T){
	str := "0123456"
	end := false
	for ;;{
		str, end = Next(str)
		fmt.Println(str)
		if end{
			break
		}
	}
}
