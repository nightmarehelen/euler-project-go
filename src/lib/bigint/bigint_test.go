package bigint

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T){
	bigint := New("12345678901234567890")
	fmt.Println(bigint.String())
}

func TestAdd(t *testing.T){
	bigint1 := New("0")
	bigint2 := New("37107287533902102798797998220837590246510135740250")

	fmt.Println(bigint1.Add(bigint2))
}
