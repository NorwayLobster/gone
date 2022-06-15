package main

import "fmt"

/*
%v 按默认格式输出，
%+v 在%v的基础上额外输出字段名，
%#v 在%+v的基础上额外输出类型名。
*/

type Teacher struct {
	name   string
	age    int
	height float64
}

func fmt_format_demo() {
	t := &Teacher{
		// t := Teacher{
		name:   "chengwang",
		age:    1,
		height: 2.0,
	}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
}
