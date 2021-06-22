package main

import "fmt"

func main() {
	//字符串的修改
	s2 := "白萝卜"
	s3 := []rune(s2)        //把字符串强制转换成了一个rune切片
	s3[0] = '红'             //修改rune切片下标第一个字符
	fmt.Println(string(s3)) //把rune切片强制转换成字符串
	c1 := "红"
	c2 := '红' //rune(int32)
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
	c3 := "H"
	c4 := byte('H')
	fmt.Printf("c3:%T c4:%T\n", c3, c4)
	//一个字符串里面有一个或多个字符，UTF-8类型的字符用rune()--int32
	//字母类型的字符用byte()--uint8
	//类型转换
	//bool类型是不能转换的
	//使用math.sqrt()函数，接收float64类型的参数。
}
