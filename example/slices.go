// 数组
package example

import (
	"fmt"
	"sort"
)

func init() {
	fmt.Println("this is example -> slices(数组)")
}

func Sl1() {
	fmt.Println("this is example -> slices -> S1")
}

func Sl2() {
	//
	slice1 := make([]int, 0, 10)
	// load the slice, cap(slice1) is 10:
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	//
	from := []int{1, 2, 3}
	to := make([]int, 10)
	// 切片复制->相同类型
	n := copy(to, from)
	fmt.Println(to)
	fmt.Printf("Copied %d elements\n", n) // n == 3
	//
	sl3 := []int{1, 2, 3}
	// 切片追加->相同类型
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
	//
	str := "going"
	// 从字符串生成字节切片
	c1 := []byte(str)
	// 拷贝方式生成
	c2 := copy(c1, str)
	//
	fmt.Println(str[0:2], c1, c2)
	//
	sl1 := []int{1, 5, 2, 8, 2, 66, 23, 99}
	// 切片排序
	sort.Ints(sl1)
	fmt.Println(sl1)
	// 切片搜索->索引
	i := sort.SearchInts(sl1, 6)
	fmt.Println(sl1[i])
}
