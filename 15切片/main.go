package main

import "fmt"

func main() {
	//	切片(slice)
	var s []int     //定义一个存放int类型的元素的切片
	var s1 []string //定义一个存放string类型的元素的切片
	fmt.Println(s, s1)
	//	初始化
	s = []int{1, 2, 3}
	s1 = []string{"泰戈尔", "真的", "帅啊"}
	fmt.Println(s, s1)
	fmt.Printf("len(s):%d cap(s):%d\n", len(s), cap(s))
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))

	//	由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7}
	a2 := a1[0:4] //基于一个数组切割，从一个位置开始，到结束位置的前一个。左包含右不包含，(左闭右开)
	fmt.Println(a2)
	a3 := a1[:4]
	a4 := a1[3:]
	a5 := a1[:]
	fmt.Println(a3, a4, a5)
	//	切片的容量是指底层数组的容量
	fmt.Printf("len(a3):%d cap(a3):%d\n", len(a3), cap(a3))
	//  底层数组从切片的第一个元素到最后的元素的数量
	fmt.Printf("len(a4):%d cap(a4):%d\n", len(a4), cap(a4))
	//切片指向一个底层的数组
	//切片的长度就是它元素的个数
	//切片的容量是底层数组从切片的第一个元素到最后一个元素的数量
}
