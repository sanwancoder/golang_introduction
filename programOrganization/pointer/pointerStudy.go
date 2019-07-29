package main

import "fmt"

func main() {
	x := 1
	p := &x      // p是一个int类型指针，指向x
	fmt.Println(x)
	*p = 100    //修改指针中的值 等价于   x= 100
	fmt.Println(x)

	//任何类型的指针的零值都是nil  指针之间也是可以进行相等测试的，只有它们指向同一个变量或者都是nil时才相等
	var i,j int
	fmt.Println(&i==&i,&i==&j,&i==nil)  //true false false


	//局部变量地址被返回后依然有效 以为指针依然引用着这个变量
	ip := f()
	fmt.Println(ip)   //输出内存地址
	fmt.Println(*ip)  //1
	*ip = 100
	fmt.Println(*ip)  //100
	fmt.Println(f()==f())  //false

}

//函数返回局部变量的地址
func f() *int{
	v := 1
	return &v; //返回地址
}
