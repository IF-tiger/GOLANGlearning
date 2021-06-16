package main

import "fmt"

// ！！！ 核心：iota在const关键字出现的时候会被重置为0，const中每新增一行常量声明iota将会自动计数一次，
const (
	s1 = iota //0
	s2        //1
	s3        //2
)
const (
	b1 = iota //0
	b2        //1
	_         //2
	b3        //3
)
const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)

//多个变量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 //d1=1,d2=2,
	d3, d4 = iota + 1, iota + 2 //d3=2,d4=3
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
}
