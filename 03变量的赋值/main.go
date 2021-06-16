package main

import "fmt"

//同一个作用域不能重复声明同名得变量

func main() {
	//1.声明变量的同时并赋值
	var s1 string = "泰gie尔"
	fmt.Println(s1)
	//2.类型推导(根据值去判断该变量是什么类型)
	var s2 = "20"
	fmt.Println(s2)
	//3.简短变量声明,只能在函数内使用
	s3 := "泰戈尔"
	fmt.Println(s3)
}
