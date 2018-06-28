// 结构体玩法
package example

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age  int
}

type File struct {
	fd   int    // 文件描述符
	name string // 文件名
}

func init() {
	fmt.Println("this is example -> structs(结构体)")
}

func St1() {

	fmt.Println("this is example -> structs -> St1")

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}

type m map[string]string

type u struct {
	uid  int    "ID"
	name string "姓名"
	age  int    "年龄"
}

func St2() {
	// struct
	u1 := new(u)
	u1.uid = 1
	u1.name = "fy"
	u1.age = 11
	fmt.Println(u1)
	// map
	m1 := make(m)
	m1["sasa"] = "232"
	fmt.Println(m1)
	// 读取tag
	tt := u{1, "alex", 18}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}

	// new 	=> 值类型		=>	指向 nil 的指针，它尚未被分配内存	(*T)
	// make => 引用类型		=>	slice, map , channel			(T)
}

// 读取tag
func refTag(tt u, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

// 内联结构体
type innerS struct {
	in1, in2 int
}

type outerS struct {
	a, b, c int
	innerS  // 取内联结构
}

func (o *outerS) AddToParam(param int) int {
	return o.a + o.b + param
}

func St3() {
	// 赋值
	outer := new(outerS)
	outer.a = 60
	outer.b = 6
	outer.c = 7
	outer.in1 = 5
	outer.in2 = 10
	fmt.Printf("outer.a is: %d\n", outer.a)
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	// 使用结构体字面量
	outer2 := outerS{6, 60, 7, innerS{5, 10}}
	fmt.Printf("outer2 is: %d\n", outer2)
	fmt.Printf("outer2.a is: %d\n", outer2.a)
	fmt.Printf("outer2.b is: %d\n", outer2.b)
	fmt.Printf("outer2.c is: %f\n", outer2.c)
	fmt.Printf("outer2.in1 is: %d\n", outer2.in1)
	fmt.Printf("outer2.in2 is: %d\n", outer2.in2)

	// 通过函数运算
	fmt.Printf("Add to the param: %d\n", outer2.AddToParam(20))

}
