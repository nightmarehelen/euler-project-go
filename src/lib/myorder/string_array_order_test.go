package myorder

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T){
	a := "12397617231239"
	o := StringOrder{a}

	fmt.Printf(o.Sort().Value)
}