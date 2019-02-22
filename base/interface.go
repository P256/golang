package main

import "fmt"

type Class struct {
	ClassId   int
	ClassName string
}

type Student struct {
	Name string
	Sex  string
	Age  int
	Class
}

func main() {
	s1 := &Student{Name: "张三", Sex: "男", Age: 66}
	fmt.Println(s1)
}
