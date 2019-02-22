package main

import "fmt"

func do(a int, b int) int {
	return a / b
}

func main() {
	//
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	do(5, 0) //异常
	fmt.Println("...")
}
