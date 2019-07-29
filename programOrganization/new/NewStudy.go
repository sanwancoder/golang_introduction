package main

import "fmt"

func main() {
	//语法糖
	//另一个创建变量的方法是调用用内建的new函数。表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为 *T
	p := new(int)  // new是一个预定义的函数 而非关键字
	fmt.Println(p)
	fmt.Println(*p)
	*p = 100
	fmt.Println(*p)
}
