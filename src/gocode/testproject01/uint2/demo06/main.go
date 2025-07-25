package main

import "fmt"

func main() {
	//转义字符
	fmt.Println("aaa\nbbb")
	//退格符
	fmt.Println("aaa\bbbb")
	//\r 光标回到行首，覆盖掉之前的字符
	fmt.Println("aaaaaaaa\rbbb")

	//\t 制表符

}
