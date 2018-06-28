package example

import (
	"fmt"
)

func init() {
	fmt.Println("this is example -> maps(字典)")
}

func M1() {
	fmt.Println("this is example -> maps -> M1")
}

func M2() {
	//
	m1 := make(map[string]string)
	m1["id"] = "1"
	m1["name"] = "小王"
	m1["age"] = "27"
	m1["sex"] = "男"
	/*for k, v := range m1 {
		fmt.Println("Map item: Capital of", k, "is", v)
	}
	for key := range m1 {
		fmt.Println("Map item: Capital of", key, "is", m1[key])
	}*/
	fmt.Println(m1)
	//
	m2 := make(map[string]string, 6)
	fmt.Println(m2)
}
