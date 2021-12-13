package main

import (
	"fmt"
	"log"
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

	now, err := ToFormatTime(LayOut)
	if err != nil {
		log.Fatal("ToFormatTime:", err)
		return
	}
	// now := time.Now()
	fmt.Printf("now: %s\n", now)
	fmt.Printf("year: %v\n", now.Year())
	h, m, s := now.Clock()

	fmt.Printf("h: %v\n", h)
	fmt.Printf("m: %v\n", m)
	fmt.Printf("s: %v\n", s)
	// fmt.Printf("clock: %s\n\n", )
	fmt.Printf("hour: %v\n", now.Hour())
	fmt.Printf("minute: %v\n", now.Minute())
	fmt.Printf("second: %v\n", now.Second())
	fmt.Println()
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

const LayOut string = "2006-01-02 15:04:05"

func ToFormatTime(strTime string) (time.Time, error) {
	timeStart, err := time.ParseInLocation(LayOut, strTime, time.Local)
	if err != nil {
		fmt.Println("Parse time failed.:", timeStart, ",err:", err)
		return time.Time{}, fmt.Errorf("Parse time failed. strTime:%s,err:%s", timeStart, err)
	}

	return timeStart, err
	// return timeStart.Unix(), err
}
