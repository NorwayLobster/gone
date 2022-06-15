/*
 * @Date: 2022-03-08 20:05:39
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-03-11 16:42:24
 * @FilePath: /gone/prometheus_sdk_demo.go
 */

package main

import (
	"fmt"
	"time"

	sdk "git.bilibili.co/gserver/prometheus_sdk"
)

// BiLog: http://10.24.10.37:2112/
// BiLog的数据提交地址（TCP）： 10.24.10.37:25555
// Prometheus: http://10.24.10.37:9090/
// grafana地址：http://10.24.10.37:3000/

func prometheus_sdk() {
	if !sdk.Init("10.24.10.37:25555", nil, 100) {
		return
	}

	// for i := 0; i < 10; i++ {
	// for {
	start := time.Now().UnixNano()
	fmt.Println("hello world")
	end := time.Now().UnixNano()
	sdk.Update_FUNCTION_EXECUTION("gone", "prometheus_sdk", (end-start)/1e6)
	// }
	select {}
}
