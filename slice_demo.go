/*
 * @Date: 2022-02-10 20:06:15
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-02-10 20:09:35
 * @FilePath: /gone/slice_demo.go
 */
package main

import "fmt"

func slice_delete_demo() {

	myNum := []int{10, 20, 30, 40, 50}
	len1 := len(myNum)
	lastIndex := len1 - 1
	fmt.Println(myNum)
	// 创建一个新切片
	// 其长度为 2 个元素，容量为 4 个元素
	myNum = append(myNum[:lastIndex], myNum[lastIndex+1:]...)
	fmt.Println(myNum)
	// newNum := myNum[1:3]
}
