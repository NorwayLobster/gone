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
