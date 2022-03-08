/*
 * @Date: 2022-03-08 18:26:41
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-03-08 18:32:09
 * @FilePath: /gone/type_assertion_demo.go
 */
package main

import (
	"errors"
	"fmt"
	"strconv"
)

func type_assertion_demo() {
	// type_assertion_demo0()
	// type_assertion_demo1()
	type_assertion_demo2()
	// type_assertion_demo3()
}

func type_assertion_demo0() {
	var x interface{}
	// x = 10
	x = "hello"
	value := x.(int)
	fmt.Println(value)
}

//https://blog.csdn.net/HaoDaWang/article/details/80748106
//x.(T): x必须为接口类型, T为具体类型或接口类型
//case1: 断言类型T为具体类型
func type_assertion_demo1() {
	var x interface{}
	x = "hello"
	// x = 10
	if value, ok := x.(int); ok != false {
		fmt.Println(value, ",", ok)
	} else {

		fmt.Println(value, ",", ok)
	}
}

func type_assertion_demo2() {
	// var str int32 = 32
	// str := 32
	str := "hello"
	err := VarType(str)
	if err != nil {
		fmt.Println("error: %w", err)
	}
}

func VarType(x interface{}) (err error) {
	switch t := x.(type) {
	case string:
		fmt.Println("string:", t) //add your operations
	case int32:
		fmt.Println("int32:", t)
	default:
		fmt.Println("other:", t)
		return errors.New("no this type")
	}
	return nil
}

//demo3
type I interface {
	Get() int
	Put(int)
}

type P interface {
	Print()
}

//定义结构体，实现接口I
type S struct {
	i int
}

func (p *S) Get() int {
	return p.i
}
func (p *S) Put(v int) {
	p.i = v
}
func (p *S) Print() {
	fmt.Println("interface p:" + strconv.Itoa(p.i))
}

//使用类型断言
//case2: 断言类型T为接口类型
func GetInt(some interface{}) int {
	if sp, ok := some.(P); ok { // 此处断言some这个接口后面隐藏的变量实现了接口P 从而调用了. P接口中的函数Print.
		sp.Print()
	}

	// return some.Get() //error
	return some.(I).Get()
}

func type_assertion_demo3() {
	s := &S{i: 5}
	// a := GetInt(s)
	fmt.Println(GetInt(s))
}
