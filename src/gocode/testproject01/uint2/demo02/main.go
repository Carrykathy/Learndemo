package main

import "fmt"

// 全局变量
var name string = "Alice"

// 一次性声明多个变量
var (
	age  int    = 25
	city string = "Beijing"
)

func main() {
	// 局部变量
	//	第一种方式，声明并赋值
	var num int = 19
	fmt.Println(num)

	//	第二种方式，声明但不赋值
	var num2 int
	fmt.Println(num2)

	//	第三种方式，声明但不写类型，由编译器自动推导
	var num3 = 10
	fmt.Println(num3)

	//	第四种方式，短变量声明
	sex := "female"
	fmt.Println(sex)

	fmt.Println("------------")
	//	声明多个变量
	var a, b, c int
	fmt.Println(a, b, c)

	//	声明多个变量并赋值
	var a1, b1, c1 = 1, 2, 3
	fmt.Println(a1, b1, c1)

	//	声明多个变量并赋值, 省略类型
	n1, n2 := 10, 100.4
	fmt.Println(n1, n2)

	fmt.Println(name)

	fmt.Println(age)
	fmt.Println(city)

}
