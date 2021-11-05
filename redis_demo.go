package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

func redis_demo() {

	rand.Seed(time.Now().UnixNano())

	addr := "127.0.0.1:6379"
	index := 6
	rawURL := fmt.Sprintf("redis://%s/%d", addr, index)
	// pwd:=12356
	// conn, err := redis.DialURL(rawURL, redis.DialPassword(pwd))
	conn, err := redis.DialURL(rawURL)
	if err != nil {
		fmt.Println("redis.DialURL failed, err:", err)
	}
	defer conn.Close()

	for i := 0; i < 10; i++ {
		var uid uint64 = uint64(i)
		// value := strconv.FormatUint(uid, 10)
		// fmt.Printf("uid: %s\n", value)

		timestamp := time.Now().Unix()
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Printf("timestamp: %v\n", uint32(timestamp))
		var level uint32 = uint32(rand.Int31n(55))
		fmt.Printf("level: %d\n", level)
		score := generate1(uint32(99)-level, uint32(timestamp))
		fmt.Printf("score: %d\n", score)
		_, err = conn.Do("ZADD", "levelall_test2", score, uid)
		// _, err = conn.Do("ZADD", "levelall", strconv.FormatUint(score, 10), value)
		if err != nil {
			fmt.Printf("ZADD %s\n", err)
			return
		}

	}

	// values, err := redis.StringMap(conn.Do("ZRANGE", "zset", 0, 100, "WITHSCORES"))
	values, err := redis.Strings(conn.Do("ZRANGE", "levelall_test2", 0, 10, "WITHSCORES"))
	if err != nil {
		fmt.Printf("ZRANGE %s\n", err)
		return
	}
	for i := 0; i < len(values); i += 2 {
		rankid := i >> 1
		uid, err := strconv.ParseUint(values[i], 10, 64)
		if err != nil {
			fmt.Printf("ParseUint %s\n", err)
			return
		}
		score, err := strconv.ParseUint(values[i+1], 10, 64)
		if err != nil {
			fmt.Printf("ParseUint %s\n", err)
			return
		}
		// fmt.Printf("after generate, score: %d\n", score)
		a, b := degenerate1(score)
		level := uint32(99) - a
		// fmt.Printf("level: %d\n", level)
		// fmt.Printf("time: %d\n", b)
		fmt.Printf("rankid: %d,uid:%d,level: %v,time: %v, score: %v\n", rankid, uid, level, b, score)
	}
}
