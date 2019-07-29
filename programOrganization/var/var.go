package main

import "fmt"

/**
 var 变量名字 类型 = 表达式
   如果表达式省略，零值进行初始化 数值型--0 布尔型--false 字符串类型---"" 接口或者引用类型---nil 数组/结构体等聚合类型---null
 */
func main()  {
	var s string
	var iv interface{}
	var i,j,k int
	var b,f,str = true,2.3,"four"
	fmt.Println("==========================================")
	fmt.Println(s)
	fmt.Println(iv)
	fmt.Printf("%g   %g   %g \n",i,j,k)
	fmt.Printf("%g   %g   %g \n",b,f,str)
	fmt.Println("==========================================")
	i = 100
	j = 10
	fmt.Printf("%g %g",i,j)
	// := 用在变量定义时候   不能用在已有变量
	// = 表示变量赋值
	i,j = j,i
	ii := 1000
	fmt.Printf("%g %g \n",i,j)
	fmt.Printf("%g \n",ii)

}
