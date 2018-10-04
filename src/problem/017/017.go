package main

import (
	"fmt"
	"strings"
)

/**
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19
letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?


NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115
(one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
 */
var dict = make(map[int]string)
func main(){
	dict[1] = "one"
	dict[2] = "two"
	dict[3] = "three"
	dict[4] = "four"
	dict[5] = "five"
	dict[6] = "six"
	dict[7] = "seven"
	dict[8] = "eight"
	dict[9] = "nine"
	dict[10] = "ten"
	dict[11] = "eleven"
	dict[12] = "twelve"
	dict[13] = "thirteen"
	dict[14] = "fourteen"
	dict[15] = "fifteen"
	dict[16] = "sixteen"
	dict[17] = "seventeen"
	dict[18] = "eighteen"
	dict[19] = "nineteen"
	dict[20] = "twenty"
	dict[30] = "thirty"
	dict[40] = "forty"
	dict[50] = "fifty"
	dict[60] = "sixty"
	dict[70] = "seventy"
	dict[80] = "eighty"
	dict[90] = "ninety"
	dict[100] = "hundred"

	sum := 0
	for i:=1; i<=999; i++{
		words := getNumWord(i)
		fmt.Printf("%s %d\n", words, countLetters(words))
		sum = sum + countLetters(words)
	}

	fmt.Println(sum+countLetters("one thousand"))
}

func getNumWord(num int) string{
	ret := strings.Builder{}
	hundred := num / 100
	if hundred != 0{
		ret.WriteString(dict[hundred]+" hundred")
	}

	tens := num%100
	if tens != 0{
		if hundred != 0{
			if tens >= 20{
				ret.WriteString(" and "+dict[tens - tens%10]+" "+dict[tens%10])
			}else{
				ret.WriteString(" and "+dict[tens])
			}
		}else{
			if tens >= 20{
				ret.WriteString(dict[tens - tens%10]+" "+dict[tens%10])
			}else{
				ret.WriteString(dict[tens])
			}
		}
	}
	return ret.String()
}

func countLetters(str string) int{
	return len(strings.Replace(str, " ", "", -1))
}
