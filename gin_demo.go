/*
 * @Date: 2022-04-06 23:54:28
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-04-07 18:28:54
 * @FilePath: /gone/gin_demo.go
 */
package main

import (
	"net/http"

	// "time/rate"
	"golang.org/x/time/rate"

	"github.com/gin-gonic/gin"
)

//令牌桶算法: time/rate
// Golang 标准库限流器 time/rate 使用介绍: https://www.cyhone.com/articles/usage-of-golang-rate/
// https://cloud.tencent.com/developer/article/1703745
//
//漏桶算法: github.com/uber-go/ratelimit

// 三种常见的限流算法:  https://www.cnblogs.com/duanxz/p/4123068.html

func startHTTPServer() {
	router := gin.Default()
	router.Use(MaxAllowed()) //限制每秒最多允许200个请求
	// router.Use(MaxAllowed(200)) //限制每秒最多允许200个请求
	router.GET("/ping", PongHandler)
	router.GET("/rankinglist/:name/:num", RankingListHandler)
	router.Run(":7080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

//MaxAllowed 限流器
func MaxAllowed() func(c *gin.Context) {
	// func MaxAllowed(limitValue int64, tokenBucketSize int) func(c *gin.Context) {
	// limiter := utils.NewLimiter(limitValue)
	limiter := rate.NewLimiter(10, 1)

	// log.Println("limiter.SetMax:", limitValue)
	// 返回限流逻辑
	return func(c *gin.Context) {
		ok := limiter.Allow()
		// ok := limiter.AllowN(time.Now(), 1)
		if !ok {
			c.AbortWithStatus(http.StatusServiceUnavailable) //超过每秒200，就返回503错误码
			return
		}
		c.Next()
	}
}

// func main(){
//     router := gin.New()
//     router.Use(MaxAllowed(200))  //限制每秒最多允许200个请求
// }
