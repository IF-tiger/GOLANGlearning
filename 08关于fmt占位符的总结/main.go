package main

import "fmt"

func main() {
	//fmt占位符
	var n = 100
	fmt.Printf("%T\n", n) //%T查看类型
	fmt.Printf("%v\n", n) //%v查看变量的值
	fmt.Printf("%b\n", n) //%b查看二进制
	fmt.Printf("%d\n", n) //%d查看十进制
	fmt.Printf("%o\n", n) //%o查看八进制
	fmt.Printf("%x\n", n) //%x查看十六进制
	var s = "tiger"
	fmt.Printf("%s\n", s)  //%s查看字符串类型的值
	fmt.Printf("%#v\n", s) //%#查看字符串类型的值的原始格式
}
