package main

import (
	"fmt"
	"time"
)

/**
You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
 */

func main() {
	start := time.Date(1901, 1, 1, 0, 0, 0, 0, time.FixedZone("+8", 8*60*60))
	end := time.Date(2000, 12, 31, 0, 0, 0, 0, time.FixedZone("+8", 8*60*60))

	cnt := 0
	for ; start.Before(end); {
		if start.Weekday() == time.Sunday && start.Day() == 1{
			cnt = cnt +1
		}
		start = start.Add(time.Hour*24)
	}

	fmt.Printf("%d Sundays fell on the first of the month during the twentieth century!\n", cnt)
}
