package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Age  int
	Name string
}

func (s *Student) String() string {
	return fmt.Sprintf("{Name: %s, Age: %d}", s.Name, s.Age)
}

type StudentSet []*Student

func (s StudentSet) Len() int {
	return len(s)
}
func (s StudentSet) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s StudentSet) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sort_demo() {

	var s StudentSet
	s = append(s, &Student{Age: 11, Name: "yang"})
	s = append(s, &Student{Age: 10, Name: "cheng"})
	s = append(s, &Student{Age: 12, Name: "xiao"})
	fmt.Println(s)
	sort.Sort(s)
	fmt.Println(s)
}
