/*
 * @Date: 2022-03-08 18:00:50
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-03-08 19:02:31
 * @FilePath: /gone/interface_embeded_into_struct.go
 */

package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

//https://blog.csdn.net/weixin_34268310/article/details/92645136
//https://go.dev/ref/spec#Method_sets
// https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html
// 接口：一组方法的集合
// OpenCloser 接口定义两个方法 返回 error
type OpenCloser interface {
	Open() error
	Close() error
}

//type Locker interface {
//	Lock() error
//	Unlock() error
//}

type Door struct {
	open bool // 门的状态是否开启
	lock bool // 门的状态是否上锁
}

func (d *Door) Open() error {
	fmt.Println("door open...")
	d.open = true
	return nil
}

func (d *Door) Close() error {
	fmt.Println("door close...")
	d.open = false
	return nil
}

type AutoDoor struct {
	OpenCloser        // 匿名接口
	delay      int    // 延迟多长时间开启
	msg        string // 自动开启时的警报
	Sth        string
}

func (a *AutoDoor) Open() error {
	fmt.Println("Open after " + strconv.Itoa(a.delay) + " seconds")
	time.Sleep(time.Duration(a.delay) * time.Second)
	fmt.Println("Door is opening:" + a.msg)
	return nil
}

func interface_embeded_into_struct() {
	// func main() {
	door := &AutoDoor{&Door{false, false}, 3, "warning", ""}
	// door := &AutoDoor{}
	fmt.Println("---")
	door.Open()
	fmt.Println("---")
	if v, ok := door.OpenCloser.(*Door); ok { //类型断言
		fmt.Println(v)
	}
	fmt.Println("-------------------------")
	fmt.Println(door.Sth)
	fmt.Println(door.OpenCloser)
	t := reflect.TypeOf(door.OpenCloser)
	fmt.Println(t)
	v := reflect.ValueOf(door.OpenCloser)
	fmt.Println(v)

	// return

	door.OpenCloser.Open()
	fmt.Println("---")
	if v, ok := door.OpenCloser.(*Door); ok { //类型断言
		fmt.Println(v)
	}
	fmt.Println("-------------------------")

	door.Close()
	fmt.Println("---")
	if v, ok := door.OpenCloser.(*Door); ok { //类型断言
		fmt.Println(v)
	}

}
