package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num uint8 = 129
	fmt.Println(num)

	var num3 = 28
	//Printf函数的作用是：格式化，把num3的类型填充到%T中
	fmt.Printf("num3 is :%T", num3)
	fmt.Println()
	//unsafe.Sizeof函数的作用是：返回变量的字节大小
	fmt.Println(unsafe.Sizeof(num3))
	// 整型变量使用保小
}
