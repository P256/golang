package module

import (
	"fmt"
)

func init() {
	// defer 延时执行<直到上层函数返回
	fmt.Println("this is module -> module1")
	//
	/*start := time.Now()
	//fmt.Println(split(10))
	//
	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}
	//
	n := 10
	for n <= 10 && n > 0 {
		echo(n)
		n--
	}
	// 计算时间
	end := time.Now()
	time := end.Sub(start)
	fmt.Println(time)*/
}

func split(sum int) (x, y int) {
	x = sum * 4 / 10
	y = sum - x
	return
}

func echo(n int) {
	if n <= 10 {
		fmt.Println(n)
	}
}

func Mod1() {
	fmt.Println("this is module -> module1 ->mod1")
}
