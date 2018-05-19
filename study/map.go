package main

import "fmt"

// map使用
func main() {
	// make声明写法
	var map1 = make(map[string]string)
	var map2 = make(map[string]map[string]string)

	//
	map1["a"] = "map1"
	map1["b"] = "888"
	//map2["a"]["b"] = "map2"
	fmt.Println(map1, map2)
}
