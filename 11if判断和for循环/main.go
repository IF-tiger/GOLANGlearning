package main

import "fmt"

func main() {
	//if条件判断
	name := "嘉giegie"
	age := 17

	if age > 18 {
		fmt.Println(name, "网吧五连坐，开黑没赢过")
	} else {
		fmt.Println("臭"+name, "滚回去写暑假作业")
	}

	//多个条件判断
	if age > 35 {
		fmt.Println("tmd"+name, "都这么大年纪了还打游戏")
	} else if age > 20 {
		fmt.Println("今晚开黑吗"+name, "人家想和你一起打游戏")
	} else {
		fmt.Println("tmd铺盖"+name, "还在打游戏，还不去写作业")
	}

	//作用域
	//如果声明的变量在语句内的话，会形成局部的作用域，好处就是代码执行完，变量就会销毁，不会占内存，保证数据的独立性。
	if age := 19; age > 18 {
		fmt.Println("网鱼网咖欢迎您")
	} else {
		fmt.Println("回家老实写作业吧")
	}

	//for循环
	//    基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//		变种for循环1
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	//	变种for循环2
	var k = 6
	for k < 10 {
		fmt.Println(k)
		k++
	}
	//	无限循环
	for {
		fmt.Println("你电脑待会就冒烟")
	}

	//	for range循环
	s := "你好泰戈尔1"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
