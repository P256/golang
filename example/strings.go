package example

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	fmt.Println("this is example -> strings(字符串)")
}

func S1() {
	fmt.Println("this is example -> strings -> S1")
}

/**
解释字符串：
    \n：换行符
    \r：回车符
    \t：tab 键
    \u 或 \U：Unicode 字符
    \\：反斜杠自身
字符串拼接符 +
*/
func S2() {
	//
	var str string = "This is an example of a string"
	// 判断字符串前缀
	isPrefix := strings.HasPrefix(str, "this")
	// 判断字符串后缀
	isSuffix := strings.HasSuffix(str, "string")
	// 判断字符串包含关系
	isContains := strings.Contains(str, "string")
	// 判断字符串中出现的位置
	indexNum := strings.Index(str, "strings")
	// 判断字符串中最后出现的位置
	indexLastNum := strings.LastIndex(str, "i")
	// 打印输出
	fmt.Println(isPrefix, isSuffix, isContains, indexNum, indexLastNum)

	// 字符串替换
	strReplace := strings.Replace(str, "string", "str", 1) + "\n"
	// 统计字符串出现次数
	strCount := strings.Count(str, "i")
	fmt.Println(strCount)
	// 重复字符串
	strRepeat := strings.Repeat(str, 2) + "\n"
	// 字符串转小写
	strToLower := strings.ToLower(str) + "\n"
	// 字符串转大写
	strToUpper := strings.ToUpper(str) + "\n"
	// 打印输出
	fmt.Println(strReplace, strRepeat, strToLower, strToUpper)

	// 修剪字符串
	strTrim := strings.Trim(strToUpper, "THIS") + "\n"
	// 分割字符串到slice
	strFields := strings.Fields(str)
	strSplit := strings.Split(str, " ")
	// 打印输出
	fmt.Println(strTrim, strFields, strSplit)
	// 拼接slice到字符串
	strJoin := strings.Join(strFields, " ")
	// 打印输出
	fmt.Println(strJoin)
	// 从字符串中读取内容
	// 生成一个 Reader 并读取字符串中的内容，然后返回指向该 Reader 的指针
	strNewReader := strings.NewReader(str)
	// Read() 从 []byte 中读取内容。
	// ReadByte() 和 ReadRune() 从字符串中读取下一个 byte 或者 rune
	fmt.Println(strNewReader)
	// 打印输出
	fmt.Printf("系统位数: %d\n", strconv.IntSize)

	// 转型：string->integer
	i, _ := strconv.Atoi("66")
	fmt.Printf("The integer is: %d\n", i)
	// 转型：integer->string
	newStr := strconv.Itoa(i)
	// 打印输出
	fmt.Printf(newStr)
}

//Match
func Match() {
	//目标字符串
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+" //正则

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}

	re, _ := regexp.Compile(pat)
	// 将匹配到的部分替换为"##.#"
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	// 参数为函数时
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	// fmt.Println(searchIn)
	fmt.Println(str2)

}
