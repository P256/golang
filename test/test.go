package main

import "fmt"

type Person struct {
	Id   int
	Name string
	Age  int
	Sex  int
}

type IPerson interface {
	GetId() int
	GetName() string
}

type IStudent interface {
	IPerson
	GetStuId() int
	GetStuName() string
}

type Student struct {
	StuId     int
	StuName   string
	ClassName string
	Person
}

func (this *Person) GetId() int {
	return this.Id
}

func (this *Person) GetName() string {
	return this.Name
}

func main() {
	fmt.Println("hello word!")
	var p1 IPerson = &Person{1, "张三", 18, 1}
	fmt.Println(p1.GetId(), p1.GetName())
}
