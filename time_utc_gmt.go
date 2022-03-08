/*
 * @Date: 2022-02-17 11:30:46
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-02-23 17:02:06
 * @FilePath: /gone/time_utc_gmt.go
 */
package main

import (
	"fmt"
	"log"
	"time"
)

var LayOut string = "2006-01-02 15:04:05"

func time_utc_gmt() {
	// time_after_before_equal()
	timeLosAngeles := "2022-02-21 16:30:25" //北京时间
	// timeLosAngeles := "2022-02-21 17:16:25" //北京时间
	shanghaiTime, err := FormatToStdTime(timeLosAngeles)
	if err != nil {
		log.Fatal("FormatToStdTime err")
	}
	fmt.Printf("shanghai time:%+v\n\n\n", shanghaiTime)

	// time_utc_gmt_demo_Los_Angeles()
	// fmt.Println()
	// time_utc_gmt_demo_Shanghai()

	// time_utc_gmt_demo()
	// time_utc_gmt_demo1()
	// time_utc_gmt_demo2()
	// time_utc_gmt_demo3()
}

func time_after_before_equal() {
	// 本地时间是 2022-2-21 17:16:25 CST, location:Local,local:2022-02-21 17:16:25.3070965 +0800 CST, isDST:false
	// loc timestamp:1645434985
	// 蛙蛙工具: 1645434985
	targetBeijing := "2022-02-21 17:16:25"  //北京时间
	targetBeijing1 := "2022-02-21 17:16:26" //北京时间target + 1s
	target := "2022-02-21 01:16:25"         //America/Los_Angeles当地时间2022-02-21 01:16:25 对应北京时间2022-2-21 17:16:25
	timeFormat := "2006-01-02 15:04:05"
	locationLos, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		log.Fatal("LoadLocation err")
	}
	shanghai_loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal("LoadLocation err")
	}

	target_in_los, err := time.ParseInLocation(timeFormat, target, locationLos)
	// target_loc, err := time.ParseInLocation(timeFormat, target, time.Local)
	if err != nil {
		log.Fatal("ParseInLocation err")
	}

	target_in_shanghai, err := time.ParseInLocation(timeFormat, targetBeijing, shanghai_loc)
	if err != nil {
		log.Fatal("ParseInLocation err")
	}
	a := target_in_los.After(target_in_shanghai)
	b := target_in_los.Before(target_in_shanghai)
	c := target_in_los.Equal(target_in_shanghai)

	fmt.Printf("a:%+v\n", a)
	fmt.Printf("b:%+v\n", b)
	fmt.Printf("c:%+v\n", c)
	fmt.Println()

	target_in_shanghai1, err := time.ParseInLocation(timeFormat, targetBeijing1, shanghai_loc)
	if err != nil {
		log.Fatal("Parse\nInLocation err")
	}

	a = target_in_los.After(target_in_shanghai1)
	b = target_in_los.Before(target_in_shanghai1)
	c = target_in_los.Equal(target_in_shanghai1)

	fmt.Printf("a:%+v\n", a)
	fmt.Printf("b:%+v\n", b)
	fmt.Printf("c:%+v\n", c)

	fmt.Println()
	a = target_in_shanghai1.After(target_in_shanghai)
	b = target_in_shanghai1.Before(target_in_shanghai)
	c = target_in_shanghai1.Equal(target_in_shanghai)
	fmt.Printf("a:%+v\n", a)
	fmt.Printf("b:%+v\n", b)
	fmt.Printf("c:%+v\n", c)
	fmt.Println()
}

func FormatToStdTime(strTime string) (time.Time, error) {
	serviceLocation := "Asia/Tokyo"
	// serviceLocation := "America/Los_Angeles"
	service_location, err := time.LoadLocation(serviceLocation)
	if err != nil {
		fmt.Println("LoadLocation err")
		return time.Time{}, fmt.Errorf("LoadLocation,err:%s", err)
	}
	service_time, err := time.ParseInLocation(LayOut, strTime, service_location)
	if err != nil {
		fmt.Println("Parse time failed.:", service_time, ",err:", err)
		return time.Time{}, fmt.Errorf("Parse time failed. strTime:%s,err:%s", service_time, err)
	}

	serverLocation := "Asia/Singapore"
	// serverLocation := "Asia/Shanghai"
	server_location, err := time.LoadLocation(serverLocation)
	if err != nil {
		fmt.Println("LoadLocation err")
		return time.Time{}, fmt.Errorf("LoadLocation,err:%s", err)
	}

	server_time := service_time.In(server_location)
	return server_time, err
}

func time_utc_gmt_demo_Los_Angeles() {
	// 本地时间是 2022-2-21 17:16:25 CST, location:Local,local:2022-02-21 17:16:25.3070965 +0800 CST, isDST:false
	// loc timestamp:1645434985
	// 蛙蛙工具: 1645434985
	// target := "2022-02-21 17:16:25"//北京时间
	target := "2022-02-21 17:16:25" //北京时间
	// target := "2022-02-21 01:16:25" //America/Los_Angeles当地时间2022-02-21 01:16:25 对应北京时间2022-2-21 17:16:25
	timeFormat := "2006-01-02 15:04:05"
	locationLos, err := time.LoadLocation("America/Los_Angeles")
	// shanghai_loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal("LoadLocation err")
	}

	target_in_los, err := time.ParseInLocation(timeFormat, target, locationLos)
	// target_loc, err := time.ParseInLocation(timeFormat, target, shanghai_loc)
	// target_loc, err := time.ParseInLocation(timeFormat, target, time.Local)
	if err != nil {
		log.Fatal("ParseInLocation err")
	}

	fmt.Printf("tar time:%+v\n", target_in_los)
	fmt.Printf("tar timestamp:%+v\n", target_in_los.Unix())
	fmt.Printf("now timestamp:%+v\n", 1645434985)

	shanghai_loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal("LoadLocation err")
	}
	target_in_los_change_to_shanghai_time := target_in_los.In(shanghai_loc)
	fmt.Printf("target_in_los change to shanghai time:%+v\n", target_in_los_change_to_shanghai_time)
	fmt.Printf("target_in_los change to shanghai timestamp:%+v\n", target_in_los_change_to_shanghai_time.Unix())

	//step2:
}

func time_utc_gmt_demo_Shanghai() {
	// 本地时间是 2022-2-21 17:16:25 CST, location:Local,local:2022-02-21 17:16:25.3070965 +0800 CST, isDST:false
	// loc timestamp:1645434985
	// 蛙蛙工具: 1645434985
	target := "2022-02-21 17:16:25" //北京时间
	// target := "2022-02-21 01:16:25" //America/Los_Angeles当地时间2022-02-21 01:16:25 对应北京时间2022-2-21 17:16:25
	timeFormat := "2006-01-02 15:04:05"
	shanghai_loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal("LoadLocation err")
	}

	// target_in_los, err := time.ParseInLocation(timeFormat, target, locationLos)
	target_loc, err := time.ParseInLocation(timeFormat, target, shanghai_loc)
	// target_loc, err := time.ParseInLocation(timeFormat, target, time.Local)
	if err != nil {
		log.Fatal("ParseInLocation err")
	}

	fmt.Printf("tar time:%+v\n", target_loc)
	fmt.Printf("tar timestamp:%+v\n", target_loc.Unix())
	fmt.Printf("now timestamp:%+v\n", 1645434985)

	// shanghai_loc, err := time.LoadLocation("Asia/Shanghai")
	locationLos, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		log.Fatal("LoadLocation err")
	}
	target_in_shanghai_change_to_Los_Angeles_time := target_loc.In(locationLos)
	fmt.Printf("target_in_los change to shanghai time:%+v\n", target_in_shanghai_change_to_Los_Angeles_time)
	fmt.Printf("target_in_los change to shanghai timestamp:%+v\n", target_in_shanghai_change_to_Los_Angeles_time.Unix())
}

func time_utc_gmt_demo() {
	// 本地时间是 2022-2-21 17:16:25 CST, location:Local,local:2022-02-21 17:16:25.3070965 +0800 CST, isDST:false
	// loc timestamp:1645434985
	target := "2022-02-21 17:16:25"
	timeFormat := "2006-01-02 15:04:05"
	shanghai_loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal("LoadLocation err")
		// panic(err)
	}
	target_loc, err := time.ParseInLocation(timeFormat, target, shanghai_loc)
	// target_loc, err := time.ParseInLocation(timeFormat, target, time.Local)
	if err != nil {
		log.Fatal("ParseInLocation err")
	}
	fmt.Printf("tar time:%+v\n", target_loc)
	fmt.Printf("tar timestamp:%+v\n", target_loc.Unix())
	fmt.Printf("now timestamp:%+v\n", 1645434985)
}

func time_utc_gmt_demo1() {
	now := time.Now()

	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()

	zone, _ := now.UTC().Zone()
	location := now.UTC().Location()
	fmt.Printf("UTC 时间是:%+v\n", now.UTC())
	fmt.Printf("UTC 时间是:%+v\n", now.In(time.UTC))
	fmt.Printf("UTC 时间是 %d-%d-%d %02d:%02d:%02d %s, %+v\n", year, mon, day, hour, min, sec, zone, location)
	utcnowtime := time.Date(year, mon, day, hour, min, sec, 0, time.UTC)
	fmt.Printf("utc timestamp:%+v\n", utcnowtime.Unix())
	fmt.Printf("time:%+v\n", utcnowtime)

	year, mon, day = now.Date()
	hour, min, sec = now.Clock()
	local := now.Local()
	location = now.Location()
	isDST := now.IsDST()
	// location = now.Local().Location()
	zone, _ = now.Zone()
	fmt.Printf("本地时间是 %d-%d-%d %02d:%02d:%02d %s, location:%+v,local:%+v, isDST:%+v\n", year, mon, day, hour, min, sec, zone, location, local, isDST)
	fmt.Printf("loc timestamp:%+v\n", now.Unix())
	fmt.Printf("time:%+v\n", now)

}

func time_utc_gmt_demo3() {
	//字符串转time.Time
	time_str := "2014-06-24 14:30:00" //时间字符串
	// Bdate := "2014-06-24 14:30:00 MST" //时间字符串
	// timeFormat := "2006-01-02 15:04:05"
	timeFormat := "2006-01-02 15:04:05"

	//local
	t_local, err := time.ParseInLocation(timeFormat, time_str, time.Local) //t被转为本地时间的time.Time
	if err != nil {
		log.Fatal("ParseInLocation err")
	}
	fmt.Printf("t:%+v\n", t_local)

	//utc
	t_utc, err := time.Parse(timeFormat, time_str) //t被转为UTC时间的time.Time
	// t_utc, err := time.Parse("2006-01-02 15:04", Bdate) //t被转为UTC时间的time.Time
	if err != nil {
		log.Fatal("ParseInLocation err")
	}
	fmt.Printf("t:%+v\n", t_utc)

	//CET
	//MST
	//CST

	//los
	locationLos, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		log.Fatal("LoadLocation err")
		// panic(err)
	}
	t_los, _ := time.ParseInLocation(timeFormat, time_str, locationLos) //t被转为UTC时间的time.Time
	fmt.Printf("t:%+v\n", t_los)
	fmt.Printf("t_loc.unix():%v\nt_utc.unix():%v\nt_los.unix():%v\n", t_local.Unix(), t_utc.Unix(), t_los.Unix())
	// fmt.Printf("los_angeles location:%v\n", t_los)
	fmt.Println("----")
	if t_local.Unix() == t_utc.Unix() {
		fmt.Printf("t_local.unix():%v, t_utc.unix():%v, equal:%v\n", t_local.Unix(), t_utc.Unix(), true)
	} else {
		fmt.Printf("t_local.unix():%v, t_utc.unix():%v, equal:%v\n", t_local.Unix(), t_utc.Unix(), false)
	}
	fmt.Println("------------------------------------------")

	timeInUTC := time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	timeInLocal := time.Date(2018, 8, 30, 12, 0, 0, 0, time.Local)
	timeInLos := time.Date(2018, 8, 30, 12, 0, 0, 0, locationLos)
	fmt.Printf("time_utc:%v\n", timeInUTC)
	fmt.Printf("time_loc:%v\n", timeInLocal)
	fmt.Printf("time_los:%v\n", timeInLos)
	fmt.Printf("timeInLos - timeInShanghai, sub:%+v\n", timeInLos.Sub(timeInLocal))
	fmt.Printf("Los time:%+v\n", time.Now().Add(timeInLos.Sub(timeInLocal)))
	fmt.Printf("Now time:%+v\n", time.Now())
	fmt.Printf("Los time:%+v\n", time.Now().In(locationLos))
	fmt.Println(timeInUTC.In(locationLos))

	fmt.Println()
	location1, err := time.LoadLocation("Local")
	if err != nil {
		panic(err)
	}
	fmt.Printf("location:%v\n", location1)
	t1 := time.Local
	fmt.Printf("t1:%v", t1)
}

func time_utc_gmt_demo2() {
	timeFormatDay := "2006-01-02"
	nowTime := time.Now()
	endTime := nowTime
	yesterdayTime := nowTime.AddDate(0, 0, -1)
	beginTime, _ := time.ParseInLocation(timeFormatDay, yesterdayTime.Format(timeFormatDay), time.Local)
	fmt.Printf("%v\n%v\n", beginTime, endTime)
	fmt.Println(time.Since(yesterdayTime))

	fmt.Println()
	timeFormat := "2006-01-02 15:04:05"
	begin, _ := time.ParseInLocation(timeFormat, "2020-06-09 01:00:00", time.Local)
	end, _ := time.ParseInLocation(timeFormat, "2020-06-10 01:00:00", time.Local)
	fmt.Printf("%v\n%v\n", begin, end)

	fmt.Println()
	locationLos, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}

	beginLos, _ := time.ParseInLocation(timeFormat, "2020-06-09 01:00:00", locationLos)
	endLos, _ := time.ParseInLocation(timeFormat, "2020-06-09 01:00:00", locationLos)
	fmt.Printf("%v\n%v\n", beginLos, endLos)
	fmt.Println()

	beginUTC, _ := time.Parse(timeFormat, "2020-06-09 01:00:00")

	fmt.Printf("%v\n%v\n", beginUTC, beginUTC)
}
