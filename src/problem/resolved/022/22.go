package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

/**
Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
 */

 type Names struct{
 	names []string
 }

 func (n *Names) Len() int{
 	return len(n.names)
 }

 func (n *Names) Less(i, j int) bool{
 	return strings.Compare(n.names[i], n.names[j]) < 0
 }

 func (n *Names) Swap(i, j int){
 	temp := n.names[i]
 	n.names[i] = n.names[j]
 	n.names[j] = temp
 }


 func main(){
 	 content,err := ioutil.ReadFile("src/file/p022_names.txt")
 	 if err != nil{
 	 	panic("fail to open file 'src/file/p022_names.txt'")
	 }

 	 namesString := strings.Split(strings.Replace(string(content), "\"", "", -1), ",")
 	 //fmt.Println(namesString)

 	 names := Names{namesString}
	 sort.Sort(&names)

 	 fmt.Println(names)
 	 fmt.Printf("score of COLIN is %d\n", getWordScore("COLIN"))

 	 sum := 0
 	 for i:=1; i<= len(names.names);i++{
 	 	sum = sum + i*getWordScore(names.names[i-1])
	 }
	fmt.Printf("Total of all the name scores in the file is %d\n", sum)
 }

 func getWordScore(name string) int{
 	score := 0
 	for _,value := range name{
 		score = score+int(value)-'A'+1
	}
 	return score
 }
