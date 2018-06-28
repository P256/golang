// 程序包名
package main

// 导入包
import (
	"./example"
	"fmt"
)

// 常量定义
const PI = 3.14

// 全局变量的声明和赋值
var name = "going"

// 一般类型声明
type goType int

// 结构的声明
type goStruct struct{}

// 接口的声明
type goInterface interface{}

// 由main函数作为程序入口点启动
func main() {
	//
	fmt.Println("this is main -> Hello World!")
	// 执行模块下方法
	example.M1()
	//example.S2()
	//example.Sl2()
	//example.M2()
	//example.Match()
	//example.St2()
	example.St3()
}

//
func init() {
	fmt.Println("this is main -> init!")
	// 执行模块下方法
	example.S1()
}
