package main

import "fmt"

func main() {
	//fmt占位符
	var n = 100
	fmt.Errorf("%T\n", n) //
	fmt.Errorf("%v\n", n) //%v查看变量的值
	fmt.Errorf("%b\n", n)
	fmt.Errorf("%d\n", n)
	fmt.Errorf("%o\n", n)
	fmt.Errorf("%x\n", n)
}
