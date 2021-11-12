package main

import (
	"fmt"
)

//1. 嵌套数据结构的打印问题
//2. map

type Info struct {
	A string
}

type IdWeight struct {
	Id     uint32
	Weight uint32
}

func (iw *IdWeight) String() string {
	return fmt.Sprintf("{Id: %d, Weight: %d}", iw.Id, iw.Weight)
}

type LuckIdToDivinationMap map[uint32][]*IdWeight

type SignConfig struct {
	Mapping LuckIdToDivinationMap
}

func initSignConfig() error {
	var monthly SignConfig
	monthly.Mapping = make(LuckIdToDivinationMap)
	fmt.Printf("%v\n", monthly)
	fmt.Printf("%+v\n", monthly)
	// monthly := GConfigHelper.MonthlySign
	// for id, config := range tbxConfig.Sign_Divination.Datas {
	var id uint32 = 100
	// if monthly.Mapping[id] == nil {
	// monthly.Mapping[id] = []*IdWeight{{Id: id, Weight: 2}}
	// } else {
	monthly.Mapping[id] = append(monthly.Mapping[id], &IdWeight{Id: id, Weight: 12})
	// }
	// }
	// go结构体存在嵌套结构体时，使用%+v格式化输出时会出现打印指针地址的问题, 如果要实现%+v格式化输出所有的内容，可以通过实现对应结构体的String()方法
	fmt.Printf("%v\n", monthly.Mapping[id])
	fmt.Printf("%+v\n", monthly.Mapping[id])
	return nil
}

func map_demo() {
	var m map[int]*Info
	fmt.Printf("len(map): %d\n", len(m))
	// fmt.Printf("cap(map): %d\n", cap(m))
	a, ok := m[1]
	fmt.Printf("a: %v, okay: %t\n", a, ok)
	fmt.Println("a: ", a, ", okay:", ok)
	m = make(map[int]*Info)
	m[100] = &Info{A: "hello"}
	m[101] = &Info{A: "world"}
	for key, value := range m {
		fmt.Printf("key: %v, value: %+v\n", key, value)
	}

	initSignConfig()
}
