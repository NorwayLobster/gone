/*
 * @Date: 2022-02-15 20:27:14
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-03-08 18:00:57
 * @FilePath: /gone/interface_demo.go
 */
package main

import "fmt"

func (s *Student) SetName(name1 string) {
	s.Name = name1
}

func (s Student) SetName1(name1 string) {
	s.Name = name1
}

type SetStudent interface {
	SetName(a string)
	// SetName1(a string)
	// SetAge1(a int)
	SetAge(a int)
}

type SetStudent1 interface {
	SetName1(a string)
	SetAge1(a int)
}

func interface_test(i SetStudent) {
	// fmt.Printf(i)
	i.SetAge(111)
	i.SetName("yangyang")
}

func interface_test1(i SetStudent1) {
	// fmt.Printf(i)
	i.SetAge1(111)
	i.SetName1("yangyang")
}
func interface_demo() {

	s1ptr := &Student{Age: 10, Name: "cheng"}
	interface_test(s1ptr) //可以使用以ptr作为接收者实现的方法,满足的接口
	fmt.Println(s1ptr)

	s1 := Student{Age: 10, Name: "cheng"}
	interface_test1(s1)
	interface_test1(s1ptr) //ptr可以使用以同时使用以value和ptr作为接收者方法而满足的接口
	fmt.Println(s1)
	fmt.Println(s1ptr)
}
