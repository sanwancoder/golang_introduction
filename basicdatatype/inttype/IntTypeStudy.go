package main

import "fmt"

func main() {
	//无符号数往往只有在位运算或其他特殊的运算场景才会使用 如: bit集合、分析二进制文件格式或者哈希加密等操作
	var u uint8 = 255
	fmt.Println(u,u+1,u*u)  //255 0 1

	var i int8=127
	fmt.Println(i,i+1,i*i) 			//127 -128 1

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n",x)  // 00100010
	fmt.Printf("%08b\n",y)		//00000110

	//向0取整
	f := 3.1415
	ii := int(f)
	fmt.Println(f,ii)   // 3.1415 3
	f = 1.99
	fmt.Println(int(f))   // 1

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d  %[1]c  %[1]q\n",ascii)
	fmt.Printf("%d  %[1]c  %[1]q\n",unicode)
	fmt.Printf("%d         %[1]q\n",newline)

}
