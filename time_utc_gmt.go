/*
 * @Date: 2022-02-17 11:30:46
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-02-18 17:41:05
 * @FilePath: /gone/time_utc_gmt.go
 */
package main

import (
	"fmt"
	"time"
)

func time_utc_gmt() {
	now := time.Now()

	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()

	zone, _ := now.UTC().Zone()
	location := now.UTC().Location()
	fmt.Printf("UTC 时间是 %d-%d-%d %02d:%02d:%02d %s, %+v\n", year, mon, day, hour, min, sec, zone, location)

	year, mon, day = now.Date()
	hour, min, sec = now.Clock()
	local := now.Local()
	location = now.Location()
	isDST := now.IsDST()
	// location = now.Local().Location()
	zone, _ = now.Zone()
	fmt.Printf("本地时间是 %d-%d-%d %02d:%02d:%02d %s, location:%+v,local:%+v, isDST:%+v\n", year, mon, day, hour, min, sec, zone, location, local, isDST)

	// //字符串转time.Time
	// Bdate := "2014-06-24 14:30:00 MST" //时间字符串

	// t_local, err := time.ParseInLocation("2006-01-02 15:04:00 UTC", Bdate, time.Local) //t被转为本地时间的time.Time
	// if err != nil {
	// 	log.Fatal("ParseInLocation err")
	// }
	// fmt.Printf("t:%+v\n", t_local)

	// t_utc, err := time.Parse("2006-01-02 15:04", Bdate) //t被转为UTC时间的time.Time
	// if err != nil {
	// 	log.Fatal("ParseInLocation err")
	// }
	// fmt.Printf("t:%+v\n", t_utc)
	// if t_local.Unix() == t_utc.Unix() {
	// 	fmt.Printf("t_local.unix():%v, t_utc.unix():%v, equal:%v\n", t_local.Unix(), t_utc.Unix(), true)
	// } else {
	// 	fmt.Printf("t_local.unix():%v, t_utc.unix():%v, equal:%v\n", t_local.Unix(), t_utc.Unix(), false)
	// }

	// location1, err := time.LoadLocation("CST")
	// if err == nil {
	// 	log.Fatal("LoadLocation, err:", err)
	// }
	//CET
	//MST

	location1, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}
	timeInUTC := time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	fmt.Printf("location:%v\n", location1)
	fmt.Println(timeInUTC.In(location1))
	location1, err = time.LoadLocation("Local")
	if err != nil {
		panic(err)
	}
	fmt.Printf("location:%v\n", location1)
	t1 := time.Local()
	fmt.Printf("t1:%v", t1)
}
