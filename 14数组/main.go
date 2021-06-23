package main

import "fmt"

func main() {
	//	数组
	//	存放元素的容器
	//	必须指定存放元素的类型和容量(长度)
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:%T a2:%T\n", a1, a2)
	//	数组的初始化
	//	如果不初始化，默认值都是零值(布尔值:false,整型和浮点型都是0,字符串: "")
	fmt.Println(a1, a2)
	//	初始化方式1 //
	a1 = [3]bool{true, true, true}
	//	初始化方式2 //根据初始值自动推断数组的长度
	a4 := [...]int{0, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a4)
	//	初始化方式3 // 根据索引来初始化
	a5 := [5]int{0: 1, 4: 3}
	fmt.Println(a5)

	//数组的遍历
	city := [...]string{"北京", "上海", "深圳"}
	//根据索引遍历
	for i := 0; i < len(city); i++ {
		fmt.Println(city[i])
	}
	//	for range遍历
	for i, v := range city {
		fmt.Println(i, v)
	}

	//	多维数组
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)
	//	多维数组的遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//	数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
}
