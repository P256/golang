package main

import "fmt"

// 变量使用
func main() {
	var (
		var0       string
		var1              = 1.1
		var2       string = "2"
		var3, var4        = "3", "4"
	)
	var5, var6 := "5", "6"
	var7 := int(var1)
	fmt.Println(var0, var1, var2, var3, var4, var5, var6, var7)
}
