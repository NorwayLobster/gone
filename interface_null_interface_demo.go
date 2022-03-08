/*
 * @Date: 2022-03-08 18:29:35
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-03-08 18:30:40
 * @FilePath: /gone/interface_null_interface_demo.go
 */
package main

import "fmt"

func null_interface_demo() {
	// test1()
	test2()
}

func test1() {
	var x interface{} //万能数据类型, 承接动态类型
	x = 1
	fmt.Println(x)
}

//root cause: 空接口是指没有定义任何接口方法的接口。没有定义任何接口方法，意味着Go中的任意对象都可以实现空接口(因为没方法需要实现)，任意对象都可以保存到空接口实例变量中。

func test2() {
	// https://www.cnblogs.com/saryli/p/13273736.html

	testSlice := []int{11, 22, 33, 44}
	var any []interface{}

	//整体赋值
	// any = testSlice //wrong

	for _, value := range testSlice { //right:
		any = append(any, value)
	}

	fmt.Println(any)

}
