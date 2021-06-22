package main

import (
	"fmt"
	"strings"
)

func main() {
	//go语言中的字符串是用双引号包裹着的
	name := "tiger"
	fmt.Printf(name)
	//单独的字母\汉字\符号表示一个字符
	c1 := 'a'
	c2 := 'b'
	fmt.Print(c1)
	fmt.Print(c2)
	//字节: 1字节= 8Bit(8个二进制位)
	//一个字符'A' = 一个字节
	//一个utf8编码的汉字,一般占三个字节

	//字符串经典windows路径案例
	//	path := "D:\GO\src\github.com\IT-Tagore\GOLANGlearning"//如果这样直接带有反斜杠的话,会将\G \g \I \G识别成占位符
	//需要转义
	path := "\"D:\\GO\\src\\github.com\\IT-Tagore\\GOLANGlearning\""
	fmt.Println(path)
	//字符串相关操作
	//查看字符串的长度
	var str1 = "tiger"
	fmt.Println(len(str1))
	//字符串拼接
	str2 := "tiger"
	str3 := "真帅"
	str4 := str2 + str3
	fmt.Println(str4)
	fmt.Printf("%s%s", str2, str3)
	str5 := fmt.Sprintf("%s%s", str2, str3)
	fmt.Println(str5)
	//分隔
	ret := strings.Split(path, "\\")
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(str3, "威猛"))
	fmt.Println(strings.Contains(str3, "真帅"))
	//前缀
	fmt.Println(strings.HasPrefix(str4, "tiger"))
	//后缀
	fmt.Println(strings.HasSuffix(str4, "真帅"))
	//第一次出现的下标
	str6 := "abcdeb"
	fmt.Println(strings.Index(str6, "b"))
	//最后一次出现的下标
	fmt.Println(strings.LastIndex(str6, "b"))
	//拼接
	fmt.Println(strings.Join(ret, "+"))
}
