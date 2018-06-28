package main

import "fmt"

// 结构体使用
func main() {
	// 声明结构
	type Books struct {
		id    int
		title string
	}
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.id = 6495407
	Book1.title = "Go语言"

	/* book 2 描述 */
	Book2.id = 6495700
	Book2.title = "Python教程"

	/* 打印 Book1 信息 */
	fmt.Println(Book1)

	/* 打印 Book2 信息 */
	fmt.Println(Book2)
}
