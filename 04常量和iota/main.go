package main

import "fmt"

//常量被定义后将不能被再次修改
//常规声明
const Π = 3.1415926

//批量声明
const (
	statusCode = 200
	errorCode  = 401
)

//批量声明注意点,如果在批量声明的时候，某个常量如果没有赋值，那么默认将和上一行的常量的值一致。
const (
	s1 = 100
	s2
	s3
)

// ！！！ iota在const关键字出现的时候会被重置为0，const中每新增一行常量声明iota将会自动计数一次，(可理解为const语句块中的行索引)
const (
	u1 = iota //0
	u2        //1
	u3        //2
)

func main() {
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
