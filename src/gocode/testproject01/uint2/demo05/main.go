package main

import "fmt"

func main() {
	//字符类型，本质上是一个整数，可以直接参与运算，输出字符的时候，会将对应的码值做一个输出

	//定义字符类型的数据
	var ch1 byte = 'a'
	fmt.Println(ch1)
	var ch2 byte = '('
	fmt.Println(ch2)
	var ch3 byte = 'l'
	fmt.Println(ch3)

	//汉字字符使用Unicode编码，可以直接参与运算

	var ch4 byte = 'A'
	//想显示对应的字符，需要使用%c进行格式化输出
	fmt.Printf("ch4对应的具体字符为：%c\n", ch4)

}
