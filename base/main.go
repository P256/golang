package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
	Sex  string
}

func (p *person) Echo() {
	fmt.Println(p.Name)
}

func main() {
	//
	fmt.Println("hello word!")
	//
	p := &person{Name: "小米", Age: 18, Sex: "男"}

	p.Echo()

	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json marshal fail fail error:%v", err)
		return
	}
	fmt.Printf("json data:%s\n", data)

	var p1 person
	err = json.Unmarshal(data, &p1)
	if err != nil {
		fmt.Printf("json unmarshal fail fail error:%v", err)
		return
	}
	fmt.Printf("%+v\n", p1)

	str := "ABCDEFG"
	for i, j := 0, len(str); i < j; i++ {
		fmt.Println(string(str[i]))
	}

}
