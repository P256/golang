package main

import (
	"fmt"
	"strings"
)

func main() {
	//
	var mapList = make(map[string]string)
	//
	mapList["a"] = "asas"
	//
	for key, value := range mapList {
		fmt.Println(key, value)
	}
	//
	var dataStr = "666,888"
	dataSplit := strings.Split(dataStr, ",")
	//
	for i := 0; i < len(dataSplit); i++ {
		fmt.Println(dataSplit[i])
	}

}
