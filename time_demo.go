package main

import (
	"fmt"
	"time"
)

func CountTotalDay(year int, month int) int {
	var num int
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			num = 30
		} else {
			num = 31
		}
	} else { //Feb.
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			num = 29
		} else {
			num = 28
		}
	}
	return num
}

func time_demo() {
	now := time.Now()
	year := now.Day()
	month := int(now.Month())
	fmt.Printf("dayId: %d\n", now.Day())
	fmt.Printf("MonthId: %d\n", now.Month())
	fmt.Printf("Day Num: %d\n", CountTotalDay(year, month))
	fmt.Printf("Day Num: %d\n", CountTotalDay(2020, 2))
	fmt.Printf("Day Num: %d\n", CountTotalDay(2000, 2))

	fmt.Println()
	for i := 1; i <= 12; i++ {
		fmt.Printf("i:%d, Day Num: %d\n", i, CountTotalDay(2100, i))
	}
}
