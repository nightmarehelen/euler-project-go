package myorder

import "sort"

type StringOrder struct{
	Value string
}

func (s *StringOrder) Len() int{
	return len(s.Value)
}

func (s *StringOrder) Swap(i, j int){
	r := []rune(s.Value)
	temp := r[i]
	r[i] = r[j]
	r[j] = temp
	s.Value = string(r)
}

func (s *StringOrder) Less(i, j int) bool{
	return s.Value[i] < s.Value[j]
}

func (s *StringOrder) Sort() *StringOrder{
	sort.Sort(s)
	return s
}

