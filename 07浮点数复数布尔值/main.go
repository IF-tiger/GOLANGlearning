package main

import "fmt"

func main() {
	//浮点数
	//math.MaxFloat32//float32的最大值
	f1 := 1.23456
	fmt.Printf("%T\n", f1) //GO语言中的小数默认都是float64类型
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2) //显示声明float32类型
	//f1 = f2//float32类型的值不能直接赋值给float64类型的变量
	//--------------------------------
	//复数
	//复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。（实部和虚部相加等于复数的类型）
	//--------------------------------
	//布尔值
	b1 := true
	var b2 bool
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T\n", b2)
	fmt.Printf("value:%v\n", b2) //布尔值，默认值是false
}
