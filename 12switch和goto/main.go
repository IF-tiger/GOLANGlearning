package main

import "fmt"

func main() {
	//	流程控制之跳出for循环
	//	 break跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break //跳出for循环
		}
		fmt.Println(i)
	}
	fmt.Println("结束")

	//	continue停止当前循环继续下一次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("结束")

	//	switch简化大量的判断
	//传统繁琐的判断
	var hand = 5
	if hand == 1 {
		fmt.Println("大拇指")
	} else if hand == 2 {
		fmt.Println("食指")
	} else {
		fmt.Println("中指")
	}

	//	switch简化判断
	switch hand {
	case 1:
		fmt.Println("大拇指")
	case 5:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")

	}

	//	goto
	//传统模式
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			break
		}
	}
	//	goto简化
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}
