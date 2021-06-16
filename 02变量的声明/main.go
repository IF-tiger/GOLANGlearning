package main

import "fmt"

// !!! go语言中的变量必须先声明才能使用

//-------------------------
// **标准声明方式**
// var 变量名 变量类型
//var name string
//var age int
//var status bool
//-------------------------

//**批量声明方式**
var (
	name   string //字符串
	age    int    //数字类型
	status bool   //布尔类型
)

//-------------------------
// GO语言中推荐使用驼峰命名法
//var vension_name string
//var vensionName string//推荐此种变量命名规范，
//var VensionName string
//-------------------------

// ！！！ GO语言中声明的非全局变量必须使用，如果不适用将停止编译
func main() {
	name = "泰gie尔"
	age = 18
	status = true

	//%s是占位符,使用name这个变量的值去替换占位符
	fmt.Print(age)              //直接在终端中输出要打印的内容
	fmt.Printf("name:%s", name) //可以进行格式化输出打印
	fmt.Println(status)         //打印完指定的内容之后会在后面加一个换行符
}
