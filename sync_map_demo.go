package main

import (
	"fmt"
	"sync"
)

func sync_map_demo() {
	sync_map_demo_test1()
}

func sync_map_demo_test1() {
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	fmt.Println(scene.Load("london"))
	scene.Store("london", 11)
	fmt.Println(scene.Load("london"))
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

	fmt.Println("-------------------------")
	//LoadOrStore方法，获取或者保存
	//参数是一对key：value，如果该key存在且没有被标记删除则返回原先的value（不更新）和true；不存在则store，返回该value 和false
	if vv, ok := scene.LoadOrStore("london", 12); ok {
		fmt.Println(vv)
	}
	fmt.Println(scene.Load("london"))
	if vv, ok := scene.LoadOrStore("beijing", 101); !ok {
		fmt.Println(vv)
	}

	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
