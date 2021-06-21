package main

import "fmt"

func main() {
	var a1 = 101
	//%d占位符替换整型十进制的数，
	fmt.Printf("%d\n", a1)
	fmt.Printf("%o\n", a1) //%o把十进制数转换成八进制
	fmt.Printf("%b\n", a1) //%o把十进制数转换成二进制
	fmt.Printf("%x\n", a1) //%x把十进制数转换成十六进制

	//八进制----0~7
	a2 := 07777
	fmt.Printf("%d\n", a2)
	//十六进制----0~F // 一般用 0x + 0~f的数
	a3 := 0xeeee
	fmt.Printf("%d\n", a3)
	//查看变量的类型
	fmt.Printf("%T\n", a3)
	//声明int8类型,必须加具体的int()的类型,否则默认就是int类型。main
	a4 := int8(9)
	fmt.Printf("%T\n", a4)
}
