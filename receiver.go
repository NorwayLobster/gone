package main

import "fmt"

func (s *Student) SetAge(a int) {
	fmt.Println("ptr")
	s.Age = a
}

func (s Student) SetAge1(a int) {
	fmt.Println("value")
	s.Age = a
}

func receiver_demo() {
	s1 := Student{Age: 10, Name: "hello"}
	s1ptr := &s1
	s1.SetAge(100) //change
	fmt.Println(s1)
	s1ptr.SetAge1(10000) //not change
	fmt.Println(s1)
}
